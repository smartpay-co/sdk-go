package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestCreateAPayment() {
	//suite.T().Skip()
	payload := CreateAPaymentJSONRequestBody{
		Order:           OrderId(suite.orderId),
		Amount:          150,
		Currency:        CurrencyJPY,
		CancelRemainder: Ptr(PaymentCreateCancelRemainderManual),
	}
	result, err := suite.client.CreateAPaymentWithResponse(suite.ctx, payload)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.NotNil(result.JSON200.Id)
	paymentId := string(*result.JSON200.Id)

	suite.Run("TestRetrieveAPayment", func() {
		params := RetrieveAPaymentParams{}
		result, err := suite.client.RetrieveAPaymentWithResponse(suite.ctx, paymentId, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		payment, _ := ConvertToStruct[Payment](result.JSON200)
		suite.EqualValues(string(*payment.Id), paymentId)
	})

	suite.Run("TestRetrieveAPaymentExpanded", func() {
		params := RetrieveAPaymentParams{
			Expand: Ptr(RetrieveAPaymentParamsExpand(ExpandAll)),
		}
		result, err := suite.client.RetrieveAPaymentWithResponse(suite.ctx, paymentId, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		payment, _ := ConvertToStruct[PaymentExpanded](result.JSON200)
		suite.EqualValues(string(*payment.Id), paymentId)
	})

	suite.Run("TestListAllPayments", func() {
		params := ListAllPaymentsParams{}
		result, err := suite.client.ListAllPaymentsWithResponse(suite.ctx, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[Payment]((*result.JSON200.Data)[0])
		suite.NotNil(obj.Id)
	})

	suite.Run("TestListAllPaymentsExpanded", func() {
		params := ListAllPaymentsParams{
			Expand: Ptr(ListAllPaymentsParamsExpand(ExpandAll)),
		}
		result, err := suite.client.ListAllPaymentsWithResponse(suite.ctx, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[PaymentExpanded]((*result.JSON200.Data)[0])
		suite.NotNil(obj.Id)
	})

	suite.Run("TestUpdateAPayment", func() {
		params := UpdateAPaymentJSONRequestBody{
			Reference: Ptr("test_update_payment"),
		}
		result, err := suite.client.UpdateAPaymentWithResponse(suite.ctx, paymentId, params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.EqualValues("test_update_payment", *result.JSON200.Reference)
	})
}
