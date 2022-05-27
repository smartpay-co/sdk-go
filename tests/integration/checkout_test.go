package integration

import (
	"context"
	. "github.com/smartpay-co/sdk-go"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type IntegrationTestSuite struct {
	client *ClientWithResponses
	suite.Suite
	ctx               context.Context
	checkoutPayload   CreateACheckoutSessionJSONRequestBody
	checkoutSessionId string
	orderId           string
}

func TestCheckOutTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func NewCheckoutPayload() *CreateACheckoutSessionJSONRequestBody {
	checkoutPayload := CreateACheckoutSessionJSONRequestBody{
		Currency: CurrencyJPY,
		Amount:   350,
		Items: []Item{
			{Name: "レブロン 18 LOW", Amount: 250, Currency: CurrencyJPY, Quantity: 1},
		},
		ShippingInfo: ShippingInfo{
			Address:     Address{Country: AddressCountryJP, PostalCode: "123", Locality: "locality", Line1: "line1"},
			FeeAmount:   Ptr(float32(100)),
			FeeCurrency: Ptr(CurrencyJPY),
		},
		CaptureMethod: Ptr(CaptureMethodManual),
		Reference:     Ptr("order_ref_1234567"),
		SuccessUrl:    "https://smartpay.co",
		CancelUrl:     "https://smartpay.co",
	}
	return &checkoutPayload
}

func (suite *IntegrationTestSuite) SetupSuite() {
	secretKey := os.Getenv("SMARTPAY_SECRET_KEY")
	publicKey := os.Getenv("SMARTPAY_PUBLIC_KEY")
	apiBase := os.Getenv("API_BASE")

	suite.ctx = context.TODO()
	var err error
	suite.client, err = NewClientWithResponses(secretKey, publicKey, WithBaseURL(apiBase))
	if err != nil {
		panic(err)
	}

	suite.checkoutPayload = *NewCheckoutPayload()
	result, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)
	if err != nil {
		panic(err)
	}
	suite.checkoutSessionId = string(*result.JSON200.Id)
	suite.orderId = string(*result.JSON200.Order.Id)
}

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
