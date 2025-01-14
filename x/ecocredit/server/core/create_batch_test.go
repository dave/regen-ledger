package core

import (
	"context"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateBatch_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, _, addr2 := testdata.KeyTestPubAddr()

	blockTime, err := types.ParseDate("block time", "2049-01-30")
	assert.NilError(t, err)
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	start, end := time.Now(), time.Now()
	res, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "C01-001",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10",
				RetiredAmount:  "5.3",
			},
			{
				Recipient:      addr2.String(),
				TradableAmount: "2.4",
				RetiredAmount:  "3.4",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
	})
	totalTradable := "12.4"
	totalRetired := "8.7"

	// check the batch
	batch, err := s.stateStore.BatchTable().Get(s.ctx, 1)
	assert.NilError(t, err, "unexpected error: %w", err)
	assert.Equal(t, res.BatchDenom, batch.Denom)
	assert.Check(t, batch.IssuanceDate.AsTime().Equal(blockTime))

	// check the supply was set
	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, totalTradable, sup.TradableAmount, "got %s", sup.TradableAmount)
	assert.Equal(t, totalRetired, sup.RetiredAmount, "got %s", sup.RetiredAmount)
	assert.Equal(t, "0", sup.CancelledAmount, "got %s", sup.CancelledAmount)

	// check balances were allocated
	bal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, "10", bal.TradableAmount)
	assert.Equal(t, "5.3", bal.RetiredAmount)

	bal2, err := s.stateStore.BatchBalanceTable().Get(s.ctx, addr2, 1)
	assert.NilError(t, err)
	assert.Equal(t, "2.4", bal2.TradableAmount)
	assert.Equal(t, "3.4", bal2.RetiredAmount)

	// check sequence number
	seq, err := s.stateStore.BatchSequenceTable().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextSequence)
}

func TestCreateBatch_BadPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)

	start, end := time.Now(), time.Now()
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "C01-001",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10.1234567891111",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")
}

func TestCreateBatch_UnauthorizedIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		ProjectId: "C01-001",
		Issuer:    sdk.AccAddress("FooBarBaz").String(),
	})
	assert.ErrorContains(t, err, "is not an issuer for the class")
}

func TestCreateBatch_ProjectNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, err := s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		ProjectId: "none",
	})
	assert.ErrorContains(t, err, "not found")
}

func TestCreateBatch_WithOriginTx_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, _, addr2 := testdata.KeyTestPubAddr()

	blockTime, err := time.Parse("2006-01-02", "2049-01-30")
	assert.NilError(t, err)
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	start, end := time.Now(), time.Now()
	_, err = s.k.CreateBatch(s.ctx, &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "C01-001",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10",
				RetiredAmount:  "5.3",
			},
			{
				Recipient:      addr2.String(),
				TradableAmount: "2.4",
				RetiredAmount:  "3.4",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
	})
	assert.NilError(t, err)
}

func TestCreateBatch_WithOriginTx_Invalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchTestSetup(t, s.ctx, s.stateStore, s.addr)
	_, _, addr2 := testdata.KeyTestPubAddr()

	blockTime, err := time.Parse("2006-01-02", "2049-01-30")
	assert.NilError(t, err)
	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	start, end := time.Now(), time.Now()
	batch := &core.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "C01-001",
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.addr.String(),
				TradableAmount: "10",
				RetiredAmount:  "5.3",
			},
			{
				Recipient:      addr2.String(),
				TradableAmount: "2.4",
				RetiredAmount:  "3.4",
			},
		},
		Metadata:  "",
		StartDate: &start,
		EndDate:   &end,
		OriginTx: &core.OriginTx{
			Id:     "210985091248",
			Source: "Ethereum",
		},
	}

	_, err = s.k.CreateBatch(s.ctx, batch)
	assert.NilError(t, err)

	// create credit batch with same tx origin id
	_, err = s.k.CreateBatch(s.ctx, batch)
	assert.ErrorContains(t, err, "credits already issued with tx id")
}

// creates a class "C01", with a single class issuer, and a project "C01-001"
func batchTestSetup(t *testing.T, ctx context.Context, ss api.StateStore, addr sdk.AccAddress) (classId, projectId string) {
	classId, projectId = "C01", "C01-001"
	classKey, err := ss.ClassTable().InsertReturningID(ctx, &api.Class{
		Id:               classId,
		Admin:            addr,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err)
	err = ss.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
		ClassKey: classKey,
		Issuer:   addr,
	})
	assert.NilError(t, err)
	_, err = ss.ProjectTable().InsertReturningID(ctx, &api.Project{
		Id:           projectId,
		ClassKey:     classKey,
		Jurisdiction: "",
		Metadata:     "",
	})
	assert.NilError(t, err)
	return
}
