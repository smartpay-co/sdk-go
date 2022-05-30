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
		retrieveAPaymentResult, err := suite.client.RetrieveAPaymentWithResponse(suite.ctx, paymentId, &params)
		suite.Nil(err)
		suite.NotNil(retrieveAPaymentResult.Body)
		suite.NotNil(retrieveAPaymentResult.JSON200)
		payment, _ := ConvertToStruct[Payment](retrieveAPaymentResult.JSON200)
		suite.EqualValues(string(*payment.Id), paymentId)
	})

	suite.Run("TestRetrieveAPaymentExpanded", func() {
		params := RetrieveAPaymentParams{
			Expand: Ptr(RetrieveAPaymentParamsExpand(ExpandAll)),
		}
		retrieveAPaymentResult, err := suite.client.RetrieveAPaymentWithResponse(suite.ctx, paymentId, &params)
		suite.Nil(err)
		suite.NotNil(retrieveAPaymentResult.Body)
		suite.NotNil(retrieveAPaymentResult.JSON200)
		payment, _ := ConvertToStruct[PaymentExpanded](retrieveAPaymentResult.JSON200)
		suite.EqualValues(string(*payment.Id), paymentId)
	})
}
