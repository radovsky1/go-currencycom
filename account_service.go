package go_currencycom

import (
	"context"
	"net/http"
)

type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "api/v2/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Account struct {
	AffiliateID      string    `json:"affiliateId"`
	Balances         []Balance `json:"balances"`
	BuyerCommission  float32   `json:"buyerCommission"`
	CanDeposit       bool      `json:"canDeposit"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	MakerCommission  float32   `json:"makerCommission"`
	SellerCommission float32   `json:"sellerCommission"`
	TakerCommission  float32   `json:"takerCommission"`
	UpdateTime       int64     `json:"updateTime"`
	UserID           int64     `json:"userId"`
}

type Balance struct {
	AccountID          string  `json:"accountId"`
	CollateralCurrency bool    `json:"collateralCurrency"`
	Asset              string  `json:"asset"`
	Free               float64 `json:"free"`
	Locked             float64 `json:"locked"`
	Default            bool    `json:"default"`
}
