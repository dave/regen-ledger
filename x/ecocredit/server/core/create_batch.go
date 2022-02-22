package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBatch creates a new batch of credits.
// Credits in the batch must not have more decimal places than the credit type's specified precision.
func (k Keeper) CreateBatch(ctx context.Context, req *v1beta1.MsgCreateBatch) (*v1beta1.MsgCreateBatchResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	projectID := req.ProjectId

	projectInfo, err := k.stateStore.ProjectInfoStore().GetByName(ctx, projectID)
	if err != nil {
		return nil, err
	}

	classInfo, err := k.stateStore.ClassInfoStore().Get(ctx, projectInfo.ClassId)
	if err != nil {
		return nil, err
	}

	err = k.assertClassIssuer(ctx, classInfo.Id, req.Issuer)
	if err != nil {
		return nil, err
	}

	p := &ecocredit.Params{}
	k.params.GetParamSet(sdkCtx, p)
	creditType, err := k.getCreditType(classInfo.CreditType, p.CreditTypes)
	if err != nil {
		return nil, err
	}

	maxDecimalPlaces := creditType.Precision

	batchSeqNo, err := k.getBatchSeqNo(ctx, projectID)
	if err != nil {
		return nil, err
	}

	batchDenom, err := ecocredit.FormatDenom(classInfo.Name, batchSeqNo, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	rowID, err := k.stateStore.BatchInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  projectInfo.Id,
		BatchDenom: batchDenom,
		Metadata:   req.Metadata,
		StartDate:  timestamppb.New(req.StartDate.UTC()),
		EndDate:    timestamppb.New(req.EndDate.UTC()),
	})
	if err != nil {
		return nil, err
	}

	tradableSupply, retiredSupply := math.NewDecFromInt64(0), math.NewDecFromInt64(0)

	for _, issuance := range req.Issuance {
		decs, err := getNonNegativeFixedDecs(maxDecimalPlaces, issuance.TradableAmount, issuance.RetiredAmount)
		if err != nil {
			return nil, err
		}
		tradable, retired := decs[0], decs[1]

		recipient, _ := sdk.AccAddressFromBech32(issuance.Recipient)
		if !tradable.IsZero() {
			tradableSupply, err = tradableSupply.Add(tradable)
			if err != nil {
				return nil, err
			}
		}
		if !retired.IsZero() {
			retiredSupply, err = retiredSupply.Add(retired)
			if err != nil {
				return nil, err
			}
			if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventRetire{
				Retirer:    recipient.String(),
				BatchDenom: batchDenom,
				Amount:     retired.String(),
				Location:   issuance.RetirementLocation,
			}); err != nil {
				return nil, err
			}
		}
		if err = k.stateStore.BatchBalanceStore().Insert(ctx, &ecocreditv1beta1.BatchBalance{
			Address:  recipient,
			BatchId:  rowID,
			Tradable: tradable.String(),
			Retired:  retired.String(),
		}); err != nil {
			return nil, err
		}

		if err = sdkCtx.EventManager().EmitTypedEvent(&v1beta1.EventReceive{
			Recipient:      recipient.String(),
			BatchDenom:     batchDenom,
			RetiredAmount:  tradable.String(),
			TradableAmount: retired.String(),
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(gasCostPerIteration, "batch issuance")
	}

	if err = k.stateStore.BatchSupplyStore().Insert(ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         rowID,
		TradableAmount:  tradableSupply.String(),
		RetiredAmount:   retiredSupply.String(),
		CancelledAmount: math.NewDecFromInt64(0).String(),
	}); err != nil {
		return nil, err
	}

	return &v1beta1.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

func (k Keeper) getBatchSeqNo(ctx context.Context, projectID string) (uint64, error) {
	var seq uint64 = 1
	batchSeq, err := k.stateStore.BatchSequenceStore().Get(ctx, projectID)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return 0, err
		}
	} else {
		seq = batchSeq.NextBatchId
	}

	if err = k.stateStore.BatchSequenceStore().Save(ctx, &ecocreditv1beta1.BatchSequence{
		ProjectId:   projectID,
		NextBatchId: seq + 1,
	}); err != nil {
		return 0, err
	}

	return seq, err
}