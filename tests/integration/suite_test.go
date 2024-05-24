package integration

import (
	"context"
	"fmt"
	"github.com/imroc/req/v3"
	. "github.com/smartpay-co/sdk-go"
	"github.com/stretchr/testify/suite"
	"net/url"
	"os"
	"path"
	"testing"
)

type IntegrationTestSuite struct {
	client *ClientWithResponses
	suite.Suite
	ctx               context.Context
	checkoutPayload   CreateACheckoutSessionJSONRequestBody
	checkoutSessionId string
	orderId           string
	testCancelOrderId string
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func NewCheckoutPayload() *CreateACheckoutSessionJSONRequestBody {
	checkoutPayload := CreateACheckoutSessionJSONRequestBody{
		Currency: CurrencyJPY,
		Amount:   1250,
		Items: []Item{
			{Name: "レブロン 18 LOW", Amount: 1000, Currency: CurrencyJPY, Quantity: Ptr(1)},
			{Name: "discount", Amount: 100, Currency: CurrencyJPY, Kind: Ptr(LineItemKindDiscount)},
			{Name: "tax", Amount: 250, Currency: CurrencyJPY, Kind: Ptr(LineItemKindTax)},
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

	suite.ctx = context.TODO()
	var err error
	suite.client, err = NewClientWithResponses(secretKey, publicKey)
	if err != nil {
		panic(err)
	}

	suite.checkoutPayload = *NewCheckoutPayload()

	// Create a checkout session for later use, to test create payment & refund
	checkoutSessionResult, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)
	if err != nil {
		panic(err)
	}
	suite.checkoutSessionId = string(*checkoutSessionResult.JSON200.Id)
	suite.orderId = string(*checkoutSessionResult.JSON200.Order.Id)
	// Authorize the order
	err = authorizeOrder(suite.orderId)
	if err != nil {
		panic(err)
	}

	// Create another checkout session for later use, to test create payment & refund
	testCancelCheckoutSessionResult, err := suite.client.CreateACheckoutSessionWithResponse(suite.ctx, suite.checkoutPayload)
	if err != nil {
		panic(err)
	}
	suite.testCancelOrderId = string(*testCancelCheckoutSessionResult.JSON200.Order.Id)
	// Authorize the order
	err = authorizeOrder(suite.testCancelOrderId)
	if err != nil {
		panic(err)
	}
}

func authorizeOrder(orderId string) (err error) {
	apiPath, _ := url.Parse(os.Getenv("API_BASE"))
	apiPath.Path = path.Join(apiPath.Path, fmt.Sprintf("/orders/%s/authorizations", orderId))
	accessToken, err := retrieveAccessToken()
	if err != nil {
		return
	}
	//req.DevMode()
	httpClient := req.C().SetLogger(nil) //.EnableDumpAll()
	_, err = httpClient.R().SetHeader("Accept", "application/json").SetBearerAuthToken(accessToken).SetBody(`{"paymentMethod":"pm_test_visaApproved","paymentPlan":"pay_in_three"}`).Post(apiPath.String())

	return
}

func authorizeToken(tokenId string) (err error) {
	apiPath, _ := url.Parse(os.Getenv("API_BASE"))
	apiPath.Path = path.Join(apiPath.Path, fmt.Sprintf("/tokens/%s/approve", tokenId))
	accessToken, err := retrieveAccessToken()
	if err != nil {
		return
	}
	//req.DevMode()
	httpClient := req.C().SetLogger(nil) //.EnableDumpAllWithoutResponse()
	_, err = httpClient.R().SetBearerAuthToken(accessToken).Put(apiPath.String())

	return
}

func retrieveAccessToken() (accessToken string, err error) {
	apiPath, _ := url.Parse(os.Getenv("API_BASE"))
	apiPath.Path = path.Join(apiPath.Path, "/consumers/auth/login")
	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")
	//req.DevMode()
	httpClient := req.C().SetLogger(nil) //.EnableDumpAllWithoutResponse()
	var dest struct {
		AccessToken *string `json:"accessToken,omitempty"`
	}
	_, err = httpClient.R().SetResult(&dest).SetBody(fmt.Sprintf(`{"emailAddress":"%s","password":"%s"}`, username, password)).Post(apiPath.String())
	accessToken = *dest.AccessToken

	return
}
