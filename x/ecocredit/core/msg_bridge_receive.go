package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgBridgeReceive{}

// Route implements the LegacyMsg interface.
func (m MsgBridgeReceive) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBridgeReceive) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBridgeReceive) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBridgeReceive) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ServiceAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("address")
	}

	// batch validation
	if m.Batch == nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch cannot be empty")
	}
	batch := m.Batch
	if _, err := sdk.AccAddressFromBech32(batch.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("recipient")
	}
	if _, err := math.NewPositiveDecFromString(batch.Amount); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf(err.Error())
	}
	if batch.OriginTx == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx is required")
	}
	if err := batch.OriginTx.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if batch.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("start_date is required")
	}
	if batch.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("end_date is required")
	}
	if batch.StartDate.After(*batch.EndDate) {
		return sdkerrors.ErrInvalidRequest.Wrap("start_date must be a time before end_date")
	}
	if len(batch.Metadata) > MaxMetadataLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch metadata length (%d) exceeds max metadata length: %d", len(batch.Metadata), MaxMetadataLength)
	}
	if len(batch.Note) > MaxMetadataLength { // do we need to do this?
		return sdkerrors.ErrInvalidRequest.Wrapf("note length (%d) exceeds max length: %d", len(batch.Note), MaxMetadataLength)
	}

	// project validation
	if m.Project == nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("project cannot be empty")
	}
	project := m.Project
	if len(project.ReferenceId) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("reference_id is required")
	}
	if err := ValidateJurisdiction(project.Jurisdiction); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if len(project.Metadata) > MaxMetadataLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("project_metadata length (%d) exceeds max metadata length: %d", len(project.Metadata), MaxMetadataLength)
	}
	if err := ValidateClassId(project.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgBridgeReceive) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ServiceAddress)
	return []sdk.AccAddress{addr}
}