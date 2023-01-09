package go_currencycom

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
		"affiliateId": "string",
		"balances": [
			{
				"accountId": "324234",
				"asset": "BTC",
				"collateralCurrency": true,
				"default": true,
				"free": 1.234,
				"locked": 0.123
			}
		],
		"buyerCommission": 0.20,
		"canDeposit": true,
		"canTrade": true,
		"canWithdraw": true,
		"makerCommission": 0.20,
		"sellerCommission": 0.20,
		"takerCommission": 0.20,
		"updateTime": 123456789,
		"userId": 123456789
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	account, err := s.client.NewGetAccountService().Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Account{
		AffiliateID:      "string",
		BuyerCommission:  0.20,
		CanDeposit:       true,
		CanTrade:         true,
		CanWithdraw:      true,
		MakerCommission:  0.20,
		SellerCommission: 0.20,
		TakerCommission:  0.20,
		UpdateTime:       123456789,
		UserID:           123456789,
		Balances: []Balance{
			{
				AccountID:          "324234",
				Asset:              "BTC",
				CollateralCurrency: true,
				Default:            true,
				Free:               1.234,
				Locked:             0.123,
			},
		},
	}
	s.assertAccountEqual(e, account)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.AffiliateID, a.AffiliateID, "AffiliateID")
	r.Equal(e.BuyerCommission, a.BuyerCommission, "BuyerCommission")
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.MakerCommission, a.MakerCommission, "MakerCommission")
	r.Equal(e.SellerCommission, a.SellerCommission, "SellerCommission")
	r.Equal(e.TakerCommission, a.TakerCommission, "TakerCommission")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.UserID, a.UserID, "UserID")
	r.Len(a.Balances, len(e.Balances))
	for i := range e.Balances {
		s.assertBalanceEqual(&e.Balances[i], &a.Balances[i])
	}
}

func (s *accountServiceTestSuite) assertBalanceEqual(e, a *Balance) {
	r := s.r()
	r.Equal(e.AccountID, a.AccountID, "AccountID")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Free, a.Free, "Free")
	r.Equal(e.Locked, a.Locked, "Locked")
	r.Equal(e.Default, a.Default, "Default")
}
