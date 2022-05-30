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
