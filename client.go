package smartpay

import (
	"context"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// SmartPay API endpoint
	Server string

	// Secret Key for accessing the API
	secretKey string

	// Public key
	publicKey string

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
		Server:    ApiEndpoint,
		secretKey: secretKey,
		publicKey: publicKey,
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
		client.Client = &http.Client{}
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
	// ListAllCheckoutSessions request
	ListAllCheckoutSessions(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateACheckoutSession request with any body
	CreateACheckoutSessionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACheckoutSession(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveACheckoutSession request
	RetrieveACheckoutSession(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllCoupons request
	ListAllCoupons(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateACoupon request with any body
	CreateACouponWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateACoupon(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveACoupon request
	RetrieveACoupon(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateACoupon request with any body
	UpdateACouponWithBody(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateACoupon(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllOrders request
	ListAllOrders(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveAnOrder request
	RetrieveAnOrder(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CancelAnOrder request
	CancelAnOrder(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAPayment request with any body
	UpdateAPaymentWithBody(ctx context.Context, paymentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAPayment(ctx context.Context, paymentId string, body UpdateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllPayments request
	ListAllPayments(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveAPayment request
	RetrieveAPayment(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateAPayment request with any body
	CreateAPaymentWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPayment(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllPromotionCodes request
	ListAllPromotionCodes(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateAPromotionCode request with any body
	CreateAPromotionCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAPromotionCode(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveAPromotionCode request
	RetrieveAPromotionCode(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAPromotionCode request with any body
	UpdateAPromotionCodeWithBody(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAPromotionCode(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllRefunds request
	ListAllRefunds(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateARefund request with any body
	CreateARefundWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateARefund(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveARefund request
	RetrieveARefund(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateARefund request with any body
	UpdateARefundWithBody(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateARefund(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListAllWebhookEndpoints request
	ListAllWebhookEndpoints(ctx context.Context, params *ListAllWebhookEndpointsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateAWebhookEndpointWithBody request with any body
	CreateAWebhookEndpointWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAWebhookEndpoint(ctx context.Context, body CreateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RetrieveAWebhookEndpoint request
	RetrieveAWebhookEndpoint(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAWebhookEndpointWithBody request with any body
	UpdateAWebhookEndpointWithBody(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAWebhookEndpoint(ctx context.Context, webhookEndpointId string, body UpdateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

// ClientWithResponsesInterface is the interface specification for the client with responses.
type ClientWithResponsesInterface interface {
	// ListAllCheckoutSessions request
	ListAllCheckoutSessionsWithResponse(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*ListAllCheckoutSessionsResponse, error)

	// CreateACheckoutSession request with any body
	CreateACheckoutSessionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	CreateACheckoutSessionWithResponse(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error)

	// RetrieveACheckoutSession request
	RetrieveACheckoutSessionWithResponse(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*RetrieveACheckoutSessionResponse, error)

	// ListAllCoupons request
	ListAllCouponsWithResponse(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*ListAllCouponsResponse, error)

	// CreateACoupon request with any body
	CreateACouponWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error)

	CreateACouponWithResponse(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error)

	// RetrieveACoupon request
	RetrieveACouponWithResponse(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*RetrieveACouponResponse, error)

	// UpdateACoupon request with any body
	UpdateACouponWithBodyWithResponse(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error)

	UpdateACouponWithResponse(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error)

	// ListAllOrders request
	ListAllOrdersWithResponse(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*ListAllOrdersResponse, error)

	// RetrieveAnOrder request
	RetrieveAnOrderWithResponse(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*RetrieveAnOrderResponse, error)

	// CancelAnOrder request
	CancelAnOrderWithResponse(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*CancelAnOrderResponse, error)

	// UpdateAPayment request with any body
	UpdateAPaymentWithBodyWithResponse(ctx context.Context, paymentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAPaymentResponse, error)

	UpdateAPaymentWithResponse(ctx context.Context, paymentId string, body UpdateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAPaymentResponse, error)

	// ListAllPayments request
	ListAllPaymentsWithResponse(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*ListAllPaymentsResponse, error)

	// RetrieveAPayment request
	RetrieveAPaymentWithResponse(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*RetrieveAPaymentResponse, error)

	// CreateAPayment request with any body
	CreateAPaymentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error)

	CreateAPaymentWithResponse(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error)

	// ListAllPromotionCodes request
	ListAllPromotionCodesWithResponse(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*ListAllPromotionCodesResponse, error)

	// CreateAPromotionCode request with any body
	CreateAPromotionCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error)

	CreateAPromotionCodeWithResponse(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error)

	// RetrieveAPromotionCode request
	RetrieveAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*RetrieveAPromotionCodeResponse, error)

	// UpdateAPromotionCode request with any body
	UpdateAPromotionCodeWithBodyWithResponse(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error)

	UpdateAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error)

	// ListAllRefunds request
	ListAllRefundsWithResponse(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*ListAllRefundsResponse, error)

	// CreateARefund request with any body
	CreateARefundWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error)

	CreateARefundWithResponse(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error)

	// RetrieveARefund request
	RetrieveARefundWithResponse(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*RetrieveARefundResponse, error)

	// UpdateARefund request with any body
	UpdateARefundWithBodyWithResponse(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error)

	UpdateARefundWithResponse(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error)

	// ListAllWebhookEndpoints request
	ListAllWebhookEndpointsWithResponse(ctx context.Context, params *ListAllWebhookEndpointsParams, reqEditors ...RequestEditorFn) (*ListAllWebhookEndpointsResponse, error)

	// CreateAWebhookEndpoint request with any body
	CreateAWebhookEndpointWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAWebhookEndpointResponse, error)

	CreateAWebhookEndpointWithResponse(ctx context.Context, body CreateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAWebhookEndpointResponse, error)

	// RetrieveAWebhookEndpoint request
	RetrieveAWebhookEndpointWithResponse(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*RetrieveAWebhookEndpointResponse, error)

	// UpdateAWebhookEndpoint request with any body
	UpdateAWebhookEndpointWithBodyWithResponse(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAWebhookEndpointResponse, error)

	UpdateAWebhookEndpointWithResponse(ctx context.Context, webhookEndpointId string, body UpdateAWebhookEndpointJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAWebhookEndpointResponse, error)
}
