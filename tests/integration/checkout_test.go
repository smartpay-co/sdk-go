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
}

func (suite *IntegrationTestSuite) TestRetrieveACheckOutSession() {
	//suite.T().Skip()
	params := RetrieveACheckoutSessionParams{}
	result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, suite.checkoutSessionId, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	checkoutSession, _ := ConvertToStruct[CheckoutSession](result.JSON200)
	suite.Equal(string(*checkoutSession.Id), suite.checkoutSessionId)
}

func (suite *IntegrationTestSuite) TestRetrieveACheckOutSessionExpanded() {
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
}

func (suite *IntegrationTestSuite) TestListAllCheckoutSession() {
	//suite.T().Skip()
	params := ListAllCheckoutSessionsParams{}
	result, err := suite.client.ListAllCheckoutSessionsWithResponse(suite.ctx, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	checkoutSession, _ := ConvertToStruct[CheckoutSession]((*result.JSON200.Data)[0])
	suite.NotNil(checkoutSession.Id)
}

func (suite *IntegrationTestSuite) TestListAllCheckoutSessionPagination() {
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
}

func (suite *IntegrationTestSuite) TestListAllCheckoutSessionExpanded() {
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
}
