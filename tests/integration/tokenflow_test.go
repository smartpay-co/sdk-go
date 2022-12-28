package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestTokenFlow() {
	//suite.T().Skip()
	checkoutPayload := CreateACheckoutSessionForATokenJSONRequestBody{
		Mode:       ModeToken,
		SuccessUrl: "https://smartpay.co",
		CancelUrl:  "https://smartpay.co",
	}
	result, err := suite.client.CreateACheckoutSessionForATokenWithResponse(suite.ctx, checkoutPayload)
	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.NotNil(result.JSON200.Token.Id)
	tokenId := string(*result.JSON200.Token.Id)

	// Authorize the Token
	err = authorizeToken(tokenId)
	if err != nil {
		panic(err)
	}

	suite.Run("TestCreateAnOrderUsingAToken", func() {
		payload := CreateAnOrderUsingATokenJSONRequestBody{
			Currency: CurrencyJPY,
			Amount:   1250,
			Items: []Item{
				{Name: "レブロン 18 LOW", Amount: 1000, Currency: CurrencyJPY, Quantity: Ptr(1)},
				{Name: "discount", Amount: 100, Currency: CurrencyJPY, Kind: Ptr(LineItemKindDiscount)},
				{Name: "tax", Amount: 250, Currency: CurrencyJPY, Kind: Ptr(LineItemKindTax)},
			},
			ShippingInfo: ShippingInfo{
				Address:     Address{Country: AddressCountryJP, PostalCode: "123", Locality: "locality", Line1: "line1"},
				FeeAmount:   Ptr(float32(100)),
				FeeCurrency: Ptr(CurrencyJPY),
			},
			CaptureMethod: Ptr(CaptureMethodManual),
			Reference:     Ptr("order_ref_1234567"),
			Token:         TokenId(tokenId),
		}
		// Create an order using the token
		result, err := suite.client.CreateAnOrderUsingATokenWithResponse(suite.ctx, payload)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.NotNil(result.JSON200.Id)
		suite.NotNil(result.JSON200.TokenId)
		suite.EqualValues(string(*result.JSON200.TokenId), tokenId)
	})

	suite.Run("TestRetrieveAToken", func() {

	})

	suite.Run("TestDisableAToken", func() {

	})

	suite.Run("TestEnableAToken", func() {

	})

	suite.Run("TestListAllTokens", func() {

	})

	suite.Run("TestDeleteAToken", func() {

	})
}
