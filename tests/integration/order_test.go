package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestListAllOrders() {
	//suite.T().Skip()
	params := ListAllOrdersParams{}
	result, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[Order]((*result.JSON200.Data)[0])
	suite.NotNil(order.Id)
}

func (suite *IntegrationTestSuite) TestListAllOrdersExpanded() {
	//suite.T().Skip()
	params := ListAllOrdersParams{
		Expand: Ptr(ListAllOrdersParamsExpand(ExpandAll)),
	}
	result, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[Order]((*result.JSON200.Data)[0])
	suite.NotNil(order.Id)
}

func (suite *IntegrationTestSuite) TestRetrieveAnOrder() {
	//suite.T().Skip()
	params := RetrieveAnOrderParams{}
	result, err := suite.client.RetrieveAnOrderWithResponse(suite.ctx, suite.orderId, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[Order](result.JSON200)
	suite.NotNil(order.Id)
}

func (suite *IntegrationTestSuite) TestRetrieveAnOrderExpanded() {
	//suite.T().Skip()
	params := RetrieveAnOrderParams{
		Expand: Ptr(RetrieveAnOrderParamsExpand(ExpandAll)),
	}
	result, err := suite.client.RetrieveAnOrderWithResponse(suite.ctx, suite.orderId, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[OrderExpanded](result.JSON200)
	suite.NotNil(order.Id)
}

func (suite *IntegrationTestSuite) TestCancelAnOrder() {
	//suite.T().Skip()
	result, err := suite.client.CancelAnOrderWithResponse(suite.ctx, suite.testCancelOrderId)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.EqualValues(string(*result.JSON200.Status), OrderStatusCanceled)
}
