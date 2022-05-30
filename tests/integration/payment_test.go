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

}
