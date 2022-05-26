package integration

import (
	"context"
	. "github.com/smartpay-co/sdk-go"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CheckOutTestSuite struct {
	client *ClientWithResponses
	suite.Suite
	checkoutPayload CreateACheckoutSessionJSONRequestBody
	*CheckoutSessionId
	ctx context.Context
}

func TestCheckOutTestSuite(t *testing.T) {
	suite.Run(t, new(CheckOutTestSuite))
}

func (suite *CheckOutTestSuite) SetupSuite() {
	secretKey := os.Getenv("SECRET_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")
	apiBase := os.Getenv("API_BASE")

	suite.ctx = context.TODO()
	var err error
	suite.client, err = NewClientWithResponses(secretKey, publicKey, WithBaseURL(apiBase))
	if err != nil {
		panic(err)
	}

	suite.checkoutPayload = CreateACheckoutSessionJSONRequestBody{
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
	result, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)
	if err != nil {
		panic(err)
	}
	suite.CheckoutSessionId = result.JSON200.Id
}

func (suite *CheckOutTestSuite) TestCreateACheckOutSession() {
	//suite.T().Skip()
	result, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
}

func (suite *CheckOutTestSuite) TestRetrieveACheckOutSession() {
	//suite.T().Skip()
	params := RetrieveACheckoutSessionParams{}
	id := string(*suite.CheckoutSessionId)
	result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, id, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.Equal(string(*result.JSON200.(*CheckoutSession).Id), id)
}

func (suite *CheckOutTestSuite) TestRetrieveACheckOutSessionExpanded() {
	//suite.T().Skip()
	params := RetrieveACheckoutSessionParams{
		Expand: Ptr(RetrieveACheckoutSessionParamsExpand(All)),
	}
	id := string(*suite.CheckoutSessionId)
	result, err := suite.client.RetrieveACheckoutSessionWithResponse(suite.ctx, id, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.Equal(string(*result.JSON200.(*CheckoutSessionExpanded).Id), id)
}
