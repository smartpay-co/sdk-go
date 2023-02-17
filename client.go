package smartpay

import (
	"context"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	DEFAULT_RETRY_WAIT_MIN = 1 * time.Second
	DEFAULT_RETRY_WAIT_MAX = 5 * time.Second
	DEFAULT_RETRY_MAX      = 3
)

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// SmartPay API endpoint
	Server string

	// Secret Key for accessing the API
	secretKey string

	// Public key
	publicKey string

	// Retry policies
	retryMax     int           // Maximum number of retries
	retryWaitMin time.Duration // Minimum time to wait
	retryWaitMax time.Duration // Maximum time to wait

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(secretKey string, publicKey string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server:       ApiEndpoint,
		secretKey:    secretKey,
		publicKey:    publicKey,
		retryMax:     DEFAULT_RETRY_MAX,
		retryWaitMin: DEFAULT_RETRY_WAIT_MIN,
		retryWaitMax: DEFAULT_RETRY_WAIT_MAX,
	}
	// read envvar for ApiEndpoint
	envApiEndpoint := os.Getenv("SMARTPAY_API_PREFIX")
	if envApiEndpoint != "" {
		client.Server = envApiEndpoint
	}

	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		retryableClient := retryablehttp.NewClient()
		retryableClient.RetryMax = client.retryMax
		retryableClient.RetryWaitMin = client.retryWaitMin
		retryableClient.RetryWaitMax = client.retryWaitMax
		retryableClient.Logger = nil
		client.Client = retryableClient.StandardClient()
	}

	// Inject dev-lang & sdk-version
	WithRequestEditorFn(func(ctx context.Context, req *http.Request) (err error) {
		values := req.URL.Query()
		values.Add("dev-lang", "go")
		values.Add("sdk-version", Version)
		req.URL.RawQuery = values.Encode()
		return
	})(&client)

	// Inject API key header
	apiKeyProvider, apiKeyProviderErr := securityprovider.NewSecurityProviderApiKey("header", "Authorization", fmt.Sprintf("Basic %s", client.secretKey))
	if apiKeyProviderErr != nil {
		panic(apiKeyProviderErr)
	}
	WithRequestEditorFn(apiKeyProvider.Intercept)(&client)

	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// WithRetryOptions allows setting up retry policies for the default http client.
// If WithHTTPClient is invoked, this function changes nothing.
func WithRetryOptions(retryMax int, retryWaitMin time.Duration, retryWaitMax time.Duration) ClientOption {
	return func(c *Client) error {
		c.retryMax = retryMax
		c.retryWaitMin = retryWaitMin
		c.retryWaitMax = retryWaitMax
		return nil
	}
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(secretKey string, publicKey string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(secretKey, publicKey, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// The interface specification for the client.
type ClientInterface interface {
	// CheckoutSession
	ListAllCheckoutSessions(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACheckoutSessionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACheckoutSession(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACheckoutSessionForAToken(ctx context.Context, body CreateACheckoutSessionForATokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveACheckoutSession(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Coupon
	ListAllCoupons(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACouponWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACoupon(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveACoupon(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateACouponWithBody(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateACoupon(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Order
	ListAllOrders(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAnOrderUsingATokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAnOrderUsingAToken(ctx context.Context, body CreateAnOrderUsingATokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveAnOrder(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CancelAnOrder(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Payment
	UpdateAPaymentWithBody(ctx context.Context, paymentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAPayment(ctx context.Context, paymentId string, body UpdateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	ListAllPayments(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveAPayment(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPaymentWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPayment(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PromotionCode
	ListAllPromotionCodes(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPromotionCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPromotionCode(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveAPromotionCode(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAPromotionCodeWithBody(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAPromotionCode(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Refund
	ListAllRefunds(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateARefundWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateARefund(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveARefund(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateARefundWithBody(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateARefund(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// WebhookEndpoint
	ListAllWebhookEndpoints(ctx context.Context, params *ListAllWebhookEndpointsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAWebhookEndpointWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAWebhookEndpoint(ctx context.Context, body CreateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveAWebhookEndpoint(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAWebhookEndpointWithBody(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAWebhookEndpoint(ctx context.Context, webhookEndpointId string, body UpdateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	DeleteAWebhookEndpoint(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Token
	ListAllTokens(ctx context.Context, params *ListAllTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	RetrieveAToken(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	EnableAToken(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	DisableAToken(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	DeleteAToken(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

// ClientWithResponsesInterface is the interface specification for the client with responses.
type ClientWithResponsesInterface interface {
	// CheckoutSession
	ListAllCheckoutSessionsWithResponse(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*ListAllCheckoutSessionsResponse, error)

	CreateACheckoutSessionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	CreateACheckoutSessionForATokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	CreateACheckoutSessionWithResponse(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	CreateACheckoutSessionForATokenWithResponse(ctx context.Context, body CreateACheckoutSessionForATokenJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	RetrieveACheckoutSessionWithResponse(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*RetrieveACheckoutSessionResponse, error)

	// Coupon
	ListAllCouponsWithResponse(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*ListAllCouponsResponse, error)

	CreateACouponWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error)

	CreateACouponWithResponse(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error)

	RetrieveACouponWithResponse(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*RetrieveACouponResponse, error)

	UpdateACouponWithBodyWithResponse(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error)

	UpdateACouponWithResponse(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error)

	// Order
	ListAllOrdersWithResponse(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*ListAllOrdersResponse, error)

	CreateAnOrderUsingATokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAnOrderUsingATokenResponse, error)

	CreateAnOrderUsingATokenWithResponse(ctx context.Context, body CreateAnOrderUsingATokenJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAnOrderUsingATokenResponse, error)

	RetrieveAnOrderWithResponse(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*RetrieveAnOrderResponse, error)

	CancelAnOrderWithResponse(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*CancelAnOrderResponse, error)

	// Payment
	UpdateAPaymentWithBodyWithResponse(ctx context.Context, paymentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAPaymentResponse, error)

	UpdateAPaymentWithResponse(ctx context.Context, paymentId string, body UpdateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAPaymentResponse, error)

	ListAllPaymentsWithResponse(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*ListAllPaymentsResponse, error)

	RetrieveAPaymentWithResponse(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*RetrieveAPaymentResponse, error)

	CreateAPaymentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error)

	CreateAPaymentWithResponse(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error)

	// PromotionCode
	ListAllPromotionCodesWithResponse(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*ListAllPromotionCodesResponse, error)

	CreateAPromotionCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error)

	CreateAPromotionCodeWithResponse(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error)

	RetrieveAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*RetrieveAPromotionCodeResponse, error)

	UpdateAPromotionCodeWithBodyWithResponse(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error)

	UpdateAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error)

	// Refund
	ListAllRefundsWithResponse(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*ListAllRefundsResponse, error)

	CreateARefundWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error)

	CreateARefundWithResponse(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error)

	RetrieveARefundWithResponse(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*RetrieveARefundResponse, error)

	UpdateARefundWithBodyWithResponse(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error)

	UpdateARefundWithResponse(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error)

	// WebhookEndpoint
	ListAllWebhookEndpointsWithResponse(ctx context.Context, params *ListAllWebhookEndpointsParams, reqEditors ...RequestEditorFn) (*ListAllWebhookEndpointsResponse, error)

	CreateAWebhookEndpointWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAWebhookEndpointResponse, error)

	CreateAWebhookEndpointWithResponse(ctx context.Context, body CreateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAWebhookEndpointResponse, error)

	RetrieveAWebhookEndpointWithResponse(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*RetrieveAWebhookEndpointResponse, error)

	UpdateAWebhookEndpointWithBodyWithResponse(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAWebhookEndpointResponse, error)

	UpdateAWebhookEndpointWithResponse(ctx context.Context, webhookEndpointId string, body UpdateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAWebhookEndpointResponse, error)

	DeleteAWebhookEndpointWithResponse(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*DeleteAWebhookEndpointResponse, error)

	// Token
	ListAllTokensWithResponse(ctx context.Context, params *ListAllTokensParams, reqEditors ...RequestEditorFn) (*ListAllTokensResponse, error)

	RetrieveATokenWithResponse(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*RetrieveATokenResponse, error)

	EnableATokenWithResponse(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*EnableATokenResponse, error)

	DisableATokenWithResponse(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*DisableATokenResponse, error)

	DeleteATokenWithResponse(ctx context.Context, tokenId string, reqEditors ...RequestEditorFn) (*DeleteATokenResponse, error)
}
