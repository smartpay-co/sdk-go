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

	suite.Run("TestCreateARefund", func() {
		params := CreateARefundJSONRequestBody{
			Amount:   Ptr(float32(1)),
			Currency: CurrencyJPY,
			Payment:  PaymentId(paymentId),
			Reason:   RefundReasonRequestedByCustomer,
		}
		result, err := suite.client.CreateARefundWithResponse(suite.ctx, params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.NotNil(result.JSON200.Id)
		refundId := string(*result.JSON200.Id)

		suite.Run("TestRetrieveARefund", func() {
			params := RetrieveARefundParams{}
			result, err := suite.client.RetrieveARefundWithResponse(suite.ctx, refundId, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			refund, _ := ConvertToStruct[Refund](result.JSON200)
			suite.EqualValues(string(*refund.Id), refundId)
		})

		suite.Run("TestRetrieveARefundExpanded", func() {
			params := RetrieveARefundParams{
				Expand: Ptr(RetrieveARefundParamsExpand(ExpandAll)),
			}
			result, err := suite.client.RetrieveARefundWithResponse(suite.ctx, refundId, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			refund, _ := ConvertToStruct[RefundExpanded](result.JSON200)
			suite.EqualValues(string(*refund.Id), refundId)
		})

		suite.Run("TestUpdateARefund", func() {
			params := UpdateARefundJSONRequestBody{
				Reference: Ptr("test_update_refund"),
			}
			result, err := suite.client.UpdateARefundWithResponse(suite.ctx, refundId, params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			suite.EqualValues("test_update_refund", *result.JSON200.Reference)
		})

		suite.Run("TestListAllRefunds", func() {
			params := ListAllRefundsParams{}
			result, err := suite.client.ListAllRefundsWithResponse(suite.ctx, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			obj, _ := ConvertToStruct[Refund]((*result.JSON200.Data)[0])
			suite.NotNil(obj.Id)
		})

		suite.Run("TestListAllRefundsExpanded", func() {
			params := ListAllRefundsParams{
				Expand: Ptr(ListAllRefundsParamsExpand(ExpandAll)),
			}
			result, err := suite.client.ListAllRefundsWithResponse(suite.ctx, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			obj, _ := ConvertToStruct[Refund]((*result.JSON200.Data)[0])
			suite.NotNil(obj.Id)
		})
	})
}
