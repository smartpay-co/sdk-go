package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestCreateACheckoutSessionForAToken() {
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
	checkoutSessionId := string(*result.JSON200.Id)
	tokenId := string(*result.JSON200.Token.Id)

	// Authorize the Token
	err = authorizeToken(tokenId)
	if err != nil {
		panic(err)
	}

	suite.Run("TestRetrieveACheckOutSessionForToken", func() {
		params := RetrieveACheckoutSessionParams{}
		result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, checkoutSessionId, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSession](result.JSON200)
		suite.Equal(string(*checkoutSession.Token), tokenId)
	})

	suite.Run("TestRetrieveACheckOutSessionForTokenExpanded", func() {
		params := RetrieveACheckoutSessionParams{
			Expand: Ptr(RetrieveACheckoutSessionParamsExpand(ExpandAll)),
		}
		result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, checkoutSessionId, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSessionExpanded](result.JSON200)
		suite.Equal(*checkoutSession.Token.Object, "token")
		suite.Equal(string(*checkoutSession.Token.Id), tokenId)
	})

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

		suite.Run("TestRetrieveAToken", func() {
			result, err := suite.client.RetrieveATokenWithResponse(suite.ctx, tokenId)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			token, _ := ConvertToStruct[Token](result.JSON200)
			suite.EqualValues(string(*token.Id), tokenId)
		})

		suite.Run("TestDisableAToken", func() {
			result, err := suite.client.DisableATokenWithResponse(suite.ctx, tokenId)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.EqualValues(result.StatusCode(), 200)
		})

		suite.Run("TestEnableAToken", func() {
			result, err := suite.client.EnableATokenWithResponse(suite.ctx, tokenId)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.EqualValues(result.StatusCode(), 200)
		})

		suite.Run("TestListAllTokens", func() {
			params := ListAllTokensParams{}
			result, err := suite.client.ListAllTokensWithResponse(suite.ctx, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			obj, _ := ConvertToStruct[Token]((*result.JSON200.Data)[0])
			suite.NotNil(obj.Id)
		})

		suite.Run("TestDeleteAToken", func() {
			result, err := suite.client.DeleteATokenWithResponse(suite.ctx, tokenId)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.EqualValues(result.StatusCode(), 204)
		})
	})

}
