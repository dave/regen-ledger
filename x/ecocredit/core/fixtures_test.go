package core

import "github.com/regen-network/regen-ledger/types/testutil"

var (
	batchDenom     = "A00-000-00000000-00000000-000"
	batchIssuance1 = BatchIssuance{Recipient: testutil.GenAddress(), TradableAmount: "12"}
	batchIssuance2 = BatchIssuance{Recipient: testutil.GenAddress(), TradableAmount: "12", RetiredAmount: "20", RetirementJurisdiction: "CH"}
	batchIssuances = []*BatchIssuance{&batchIssuance1, &batchIssuance2}
	batchOriginTx  = OriginTx{Id: "0x1234", Source: "polygon"}
)
