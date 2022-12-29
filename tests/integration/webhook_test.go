package integration

import (
	. "github.com/smartpay-co/sdk-go"
)

func (suite *IntegrationTestSuite) TestCreateAWebhookEndpoint() {
	//suite.T().Skip()
	payload := CreateAWebhookEndpointJSONRequestBody{
		EventSubscriptions: Ptr([]EventSubscription{
			EventSubscriptionMerchantUserCreated,
			EventSubscriptionPromotionCodeUpdated,
		}),
		Url: "https://www.google.com",
	}
	result, err := suite.client.CreateAWebhookEndpointWithResponse(suite.ctx, payload)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.NotNil(result.JSON200.Id)
	obj, _ := ConvertToStruct[WebhookEndpoint](result.JSON200)
	webhookEndpointId := string(*obj.Id)

	suite.Run("TestRetrieveAWebhookEndpoint", func() {
		result, err := suite.client.RetrieveAWebhookEndpointWithResponse(suite.ctx, webhookEndpointId)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[WebhookEndpoint](result.JSON200)
		suite.EqualValues(string(*obj.Id), webhookEndpointId)
	})

	suite.Run("TestUpdateAWebhookEndpoint", func() {
		payload := UpdateAWebhookEndpointJSONRequestBody{
			Metadata: Ptr(Metadata{
				"testKey1": "testValue1",
				"testKey2": "testValue2",
			}),
		}
		result, err := suite.client.UpdateAWebhookEndpointWithResponse(suite.ctx, webhookEndpointId, payload)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[WebhookEndpoint](result.JSON200)
		suite.EqualValues(string(*obj.Id), webhookEndpointId)
	})

	suite.Run("TestListAllWebhookEndpoints", func() {
		params := ListAllWebhookEndpointsParams{}
		result, err := suite.client.ListAllWebhookEndpointsWithResponse(suite.ctx, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[WebhookEndpoint]((*result.JSON200.Data)[0])
		suite.NotNil(obj.Id)
	})

	suite.Run("TestDeleteAWebhookEndpoint", func() {
		result, err := suite.client.DeleteAWebhookEndpointWithResponse(suite.ctx, webhookEndpointId)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.EqualValues(result.StatusCode(), 204)
	})
}
