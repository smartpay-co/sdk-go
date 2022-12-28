package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestCreateACheckOutSession() {
	//suite.T().Skip()
	result, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.Run("TestCheckoutSessionUrlWithPromotionCode", func() {
		result, err := result.JSON200.Url.WithPromotionCode("promo123")
		suite.Nil(err)
		suite.Contains(result, "promo123")
	})

	suite.Run("TestRetrieveACheckOutSession", func() {
		//suite.T().Skip()
		params := RetrieveACheckoutSessionParams{}
		result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, suite.checkoutSessionId, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSession](result.JSON200)
		suite.Equal(string(*checkoutSession.Id), suite.checkoutSessionId)
	})

	suite.Run("TestRetrieveACheckOutSessionExpanded", func() {
		//suite.T().Skip()
		params := RetrieveACheckoutSessionParams{
			Expand: Ptr(RetrieveACheckoutSessionParamsExpand(ExpandAll)),
		}
		result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, suite.checkoutSessionId, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSessionExpanded](result.JSON200)
		suite.Equal(string(*checkoutSession.Id), suite.checkoutSessionId)
	})

	suite.Run("TestListAllCheckoutSession", func() {
		//suite.T().Skip()
		params := ListAllCheckoutSessionsParams{}
		result, err := suite.client.ListAllCheckoutSessionsWithResponse(suite.ctx, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSession]((*result.JSON200.Data)[0])
		suite.NotNil(checkoutSession.Id)
	})

	suite.Run("TestListAllCheckoutSessionPagination", func() {
		//suite.T().Skip()
		// Get the first page
		firstPageParams := ListAllCheckoutSessionsParams{
			MaxResults: Ptr(MaxResults(1)),
		}
		FirstPageResult, err := suite.client.ListAllCheckoutSessionsWithResponse(suite.ctx, &firstPageParams)
		suite.Nil(err)
		suite.Len(*FirstPageResult.JSON200.Data, 1)

		// Get the next page
		nextPageToken := FirstPageResult.JSON200.NextPageToken
		nextPageParams := ListAllCheckoutSessionsParams{
			MaxResults: Ptr(MaxResults(1)),
			PageToken:  Ptr(PageToken(*nextPageToken)),
		}
		nextPageResult, err := suite.client.ListAllCheckoutSessionsWithResponse(suite.ctx, &nextPageParams)
		suite.Nil(err)
		suite.Len(*nextPageResult.JSON200.Data, 1)
		suite.EqualValues(string(*nextPageResult.JSON200.PageToken), string(*nextPageToken))
	})

	suite.Run("TestListAllCheckoutSessionExpanded", func() {
		//suite.T().Skip()
		params := ListAllCheckoutSessionsParams{
			Expand: Ptr(ListAllCheckoutSessionsParamsExpand(ExpandAll)),
		}
		result, err := suite.client.ListAllCheckoutSessionsWithResponse(suite.ctx, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		checkoutSession, _ := ConvertToStruct[CheckoutSessionExpanded]((*result.JSON200.Data)[0])
		suite.NotNil(checkoutSession.Id)
	})
}
