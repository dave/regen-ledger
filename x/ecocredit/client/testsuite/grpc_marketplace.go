package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

const marketplaceRoute = "regen/ecocredit/marketplace/v1"

func (s *IntegrationTestSuite) TestQuerySellOrder() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/sell-order/%d", s.val.APIAddress, marketplaceRoute, s.sellOrderId),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/sell-orders/%d", s.val.APIAddress, marketplaceRoute, s.sellOrderId),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res marketplace.QuerySellOrderResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.SellOrder)
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrders() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/sell-orders", s.val.APIAddress, marketplaceRoute),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/sell-orders?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/sell-orders?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				marketplaceRoute,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res marketplace.QuerySellOrdersResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.SellOrders)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.SellOrders, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByBatch() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/sell-orders-by-batch/%s",
				s.val.APIAddress,
				marketplaceRoute,
				s.batchDenom,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/sell-orders-by-batch/%s?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/sell-orders/batch/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				marketplaceRoute,
				s.batchDenom,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/sell-orders/batch/%s",
				s.val.APIAddress,
				marketplaceRoute,
				s.batchDenom,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res marketplace.QuerySellOrdersByBatchResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.SellOrders)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.SellOrders, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersBySeller() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/sell-orders-by-seller/%s",
				s.val.APIAddress,
				marketplaceRoute,
				s.addr1,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/sell-orders-by-seller/%s?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/sell-orders/seller/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				marketplaceRoute,
				s.addr1,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf(
				"%s/%s/sell-orders/seller/%s",
				s.val.APIAddress,
				marketplaceRoute,
				s.addr1,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res marketplace.QuerySellOrdersBySellerResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.SellOrders)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.SellOrders, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedDenoms() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/allowed-denoms", s.val.APIAddress, marketplaceRoute),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/allowed-denoms?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/allowed-denoms?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				marketplaceRoute,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res marketplace.QueryAllowedDenomsResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.AllowedDenoms)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.AllowedDenoms, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			}
		})
	}
}