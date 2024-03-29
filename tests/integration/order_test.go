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

	suite.Run("TestListAllOrdersExpanded", func() {
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
	})

	suite.Run("TestListAllOrdersPagination", func() {
		//suite.T().Skip()
		// Get the first page
		firstPageParams := ListAllOrdersParams{
			MaxResults: Ptr(MaxResults(1)),
		}
		FirstPageResult, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &firstPageParams)
		suite.Nil(err)
		suite.Len(*FirstPageResult.JSON200.Data, 1)

		// Get the next page
		nextPageToken := FirstPageResult.JSON200.NextPageToken
		nextPageParams := ListAllOrdersParams{
			MaxResults: Ptr(MaxResults(1)),
			PageToken:  Ptr(PageToken(*nextPageToken)),
		}
		nextPageResult, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &nextPageParams)
		suite.Nil(err)
		suite.Len(*nextPageResult.JSON200.Data, 1)
		suite.EqualValues(string(*nextPageResult.JSON200.PageToken), string(*nextPageToken))
	})

	suite.Run("TestRetrieveAnOrder", func() {
		//suite.T().Skip()
		params := RetrieveAnOrderParams{}
		result, err := suite.client.RetrieveAnOrderWithResponse(suite.ctx, suite.orderId, &params)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)

		order, _ := ConvertToStruct[Order](result.JSON200)
		suite.NotNil(order.Id)
	})

	suite.Run("TestRetrieveAnOrderExpanded", func() {
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

		suite.Run("TestRetrieveAnOrderExpandedSupportsNewLineItem", func() {
			for _, item := range *order.LineItems {
				if *item.Kind == LineItemKindTax {
					suite.EqualValues(int(*item.Amount), 250)
					suite.EqualValues(*item.Currency, CurrencyJPY)
				}
				if *item.Kind == LineItemKindDiscount {
					suite.EqualValues(int(*item.Amount), 100)
					suite.EqualValues(*item.Currency, CurrencyJPY)
				}
			}
		})
	})

	suite.Run("TestCancelAnOrder", func() {
		//suite.T().Skip()
		result, err := suite.client.CancelAnOrderWithResponse(suite.ctx, suite.testCancelOrderId)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.EqualValues(string(*result.JSON200.Status), OrderStatusCanceled)
	})

}
