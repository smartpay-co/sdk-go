package smartpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

func (c *Client) ListAllCheckoutSessions(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllCheckoutSessionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateACheckoutSessionWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateACheckoutSessionRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateACheckoutSession(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateACheckoutSessionRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveACheckoutSession(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveACheckoutSessionRequest(c.Server, checkoutSessionId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListAllCoupons(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllCouponsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateACouponWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateACouponRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateACoupon(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateACouponRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveACoupon(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveACouponRequest(c.Server, couponId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateACouponWithBody(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateACouponRequestWithBody(c.Server, couponId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateACoupon(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateACouponRequest(c.Server, couponId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListAllOrders(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllOrdersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveAnOrder(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveAnOrderRequest(c.Server, orderId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CancelAnOrder(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCancelAnOrderRequest(c.Server, orderId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListAllPayments(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllPaymentsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveAPayment(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveAPaymentRequest(c.Server, paymentId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAPaymentWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAPaymentRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAPayment(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAPaymentRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListAllPromotionCodes(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllPromotionCodesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAPromotionCodeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAPromotionCodeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAPromotionCode(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAPromotionCodeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveAPromotionCode(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveAPromotionCodeRequest(c.Server, promotionCodeId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAPromotionCodeWithBody(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAPromotionCodeRequestWithBody(c.Server, promotionCodeId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAPromotionCode(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAPromotionCodeRequest(c.Server, promotionCodeId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListAllRefunds(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAllRefundsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateARefundWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateARefundRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateARefund(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateARefundRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RetrieveARefund(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRetrieveARefundRequest(c.Server, refundId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateARefundWithBody(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateARefundRequestWithBody(c.Server, refundId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateARefund(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateARefundRequest(c.Server, refundId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetV1WebhookEndpoints(ctx context.Context, params *GetV1WebhookEndpointsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetV1WebhookEndpointsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostV1WebhookEndpointsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1WebhookEndpointsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostV1WebhookEndpoints(ctx context.Context, body PostV1WebhookEndpointsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostV1WebhookEndpointsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetV1WebhookEndpointsWebhookEndpointId(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetV1WebhookEndpointsWebhookEndpointIdRequest(c.Server, webhookEndpointId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchV1WebhookEndpointsWebhookEndpointIdWithBody(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchV1WebhookEndpointsWebhookEndpointIdRequestWithBody(c.Server, webhookEndpointId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchV1WebhookEndpointsWebhookEndpointId(ctx context.Context, webhookEndpointId string, body PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchV1WebhookEndpointsWebhookEndpointIdRequest(c.Server, webhookEndpointId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListAllCheckoutSessionsRequest generates requests for ListAllCheckoutSessions
func NewListAllCheckoutSessionsRequest(server string, params *ListAllCheckoutSessionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/checkout-sessions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateACheckoutSessionRequest calls the generic CreateACheckoutSession builder with application/json body
func NewCreateACheckoutSessionRequest(server string, body CreateACheckoutSessionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateACheckoutSessionRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateACheckoutSessionRequestWithBody generates requests for CreateACheckoutSession with any type of body
func NewCreateACheckoutSessionRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/checkout-sessions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRetrieveACheckoutSessionRequest generates requests for RetrieveACheckoutSession
func NewRetrieveACheckoutSessionRequest(server string, checkoutSessionId string, params *RetrieveACheckoutSessionParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "checkoutSessionId", runtime.ParamLocationPath, checkoutSessionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/checkout-sessions/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListAllCouponsRequest generates requests for ListAllCoupons
func NewListAllCouponsRequest(server string, params *ListAllCouponsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/coupons")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()
	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateACouponRequest calls the generic CreateACoupon builder with application/json body
func NewCreateACouponRequest(server string, body CreateACouponJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateACouponRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateACouponRequestWithBody generates requests for CreateACoupon with any type of body
func NewCreateACouponRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/coupons")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRetrieveACouponRequest generates requests for RetrieveACoupon
func NewRetrieveACouponRequest(server string, couponId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "couponId", runtime.ParamLocationPath, couponId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/coupons/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateACouponRequest calls the generic UpdateACoupon builder with application/json body
func NewUpdateACouponRequest(server string, couponId string, body UpdateACouponJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateACouponRequestWithBody(server, couponId, "application/json", bodyReader)
}

// NewUpdateACouponRequestWithBody generates requests for UpdateACoupon with any type of body
func NewUpdateACouponRequestWithBody(server string, couponId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "couponId", runtime.ParamLocationPath, couponId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/coupons/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListAllOrdersRequest generates requests for ListAllOrders
func NewListAllOrdersRequest(server string, params *ListAllOrdersParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/orders")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRetrieveAnOrderRequest generates requests for RetrieveAnOrder
func NewRetrieveAnOrderRequest(server string, orderId string, params *RetrieveAnOrderParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orderId", runtime.ParamLocationPath, orderId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/orders/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCancelAnOrderRequest generates requests for CancelAnOrder
func NewCancelAnOrderRequest(server string, orderId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orderId", runtime.ParamLocationPath, orderId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/orders/%s/cancellation", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListAllPaymentsRequest generates requests for ListAllPayments
func NewListAllPaymentsRequest(server string, params *ListAllPaymentsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/payments")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRetrieveAPaymentRequest generates requests for RetrieveAnOrder
func NewRetrieveAPaymentRequest(server string, paymentId string, params *RetrieveAPaymentParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "paymentId", runtime.ParamLocationPath, paymentId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/payments/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateAPaymentRequest calls the generic CreateAPayment builder with application/json body
func NewCreateAPaymentRequest(server string, body CreateAPaymentJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAPaymentRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateAPaymentRequestWithBody generates requests for CreateAPayment with any type of body
func NewCreateAPaymentRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/payments")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListAllPromotionCodesRequest generates requests for ListAllPromotionCodes
func NewListAllPromotionCodesRequest(server string, params *ListAllPromotionCodesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/promotion-codes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateAPromotionCodeRequest calls the generic CreateAPromotionCode builder with application/json body
func NewCreateAPromotionCodeRequest(server string, body CreateAPromotionCodeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAPromotionCodeRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateAPromotionCodeRequestWithBody generates requests for CreateAPromotionCode with any type of body
func NewCreateAPromotionCodeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/promotion-codes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRetrieveAPromotionCodeRequest generates requests for RetrieveAPromotionCode
func NewRetrieveAPromotionCodeRequest(server string, promotionCodeId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "promotionCodeId", runtime.ParamLocationPath, promotionCodeId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/promotion-codes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateAPromotionCodeRequest calls the generic UpdateAPromotionCode builder with application/json body
func NewUpdateAPromotionCodeRequest(server string, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateAPromotionCodeRequestWithBody(server, promotionCodeId, "application/json", bodyReader)
}

// NewUpdateAPromotionCodeRequestWithBody generates requests for UpdateAPromotionCode with any type of body
func NewUpdateAPromotionCodeRequestWithBody(server string, promotionCodeId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "promotionCodeId", runtime.ParamLocationPath, promotionCodeId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/promotion-codes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListAllRefundsRequest generates requests for ListAllRefunds
func NewListAllRefundsRequest(server string, params *ListAllRefundsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/refunds")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateARefundRequest calls the generic CreateARefund builder with application/json body
func NewCreateARefundRequest(server string, body CreateARefundJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateARefundRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateARefundRequestWithBody generates requests for CreateARefund with any type of body
func NewCreateARefundRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/refunds")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRetrieveARefundRequest generates requests for RetrieveARefund
func NewRetrieveARefundRequest(server string, refundId string, params *RetrieveARefundParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "refundId", runtime.ParamLocationPath, refundId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/refunds/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Expand != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expand", runtime.ParamLocationQuery, *params.Expand); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateARefundRequest calls the generic UpdateARefund builder with application/json body
func NewUpdateARefundRequest(server string, refundId string, body UpdateARefundJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateARefundRequestWithBody(server, refundId, "application/json", bodyReader)
}

// NewUpdateARefundRequestWithBody generates requests for UpdateARefund with any type of body
func NewUpdateARefundRequestWithBody(server string, refundId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "refundId", runtime.ParamLocationPath, refundId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/refunds/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetV1WebhookEndpointsRequest generates requests for GetV1WebhookEndpoints
func NewGetV1WebhookEndpointsRequest(server string, params *GetV1WebhookEndpointsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/webhook-endpoints")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.MaxResults != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxResults", runtime.ParamLocationQuery, *params.MaxResults); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PageToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "pageToken", runtime.ParamLocationQuery, *params.PageToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostV1WebhookEndpointsRequest calls the generic PostV1WebhookEndpoints builder with application/json body
func NewPostV1WebhookEndpointsRequest(server string, body PostV1WebhookEndpointsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostV1WebhookEndpointsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostV1WebhookEndpointsRequestWithBody generates requests for PostV1WebhookEndpoints with any type of body
func NewPostV1WebhookEndpointsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/webhook-endpoints")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetV1WebhookEndpointsWebhookEndpointIdRequest generates requests for GetV1WebhookEndpointsWebhookEndpointId
func NewGetV1WebhookEndpointsWebhookEndpointIdRequest(server string, webhookEndpointId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "webhookEndpointId", runtime.ParamLocationPath, webhookEndpointId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/webhook-endpoints/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPatchV1WebhookEndpointsWebhookEndpointIdRequest calls the generic PatchV1WebhookEndpointsWebhookEndpointId builder with application/json body
func NewPatchV1WebhookEndpointsWebhookEndpointIdRequest(server string, webhookEndpointId string, body PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchV1WebhookEndpointsWebhookEndpointIdRequestWithBody(server, webhookEndpointId, "application/json", bodyReader)
}

// NewPatchV1WebhookEndpointsWebhookEndpointIdRequestWithBody generates requests for PatchV1WebhookEndpointsWebhookEndpointId with any type of body
func NewPatchV1WebhookEndpointsWebhookEndpointIdRequestWithBody(server string, webhookEndpointId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "webhookEndpointId", runtime.ParamLocationPath, webhookEndpointId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/webhook-endpoints/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

type ListAllCheckoutSessionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]CheckoutSessionOptions `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllCheckoutSessionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllCheckoutSessionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateACheckoutSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CheckoutSessionExpanded
	JSON400      *interface{}
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateACheckoutSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateACheckoutSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveACheckoutSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CheckoutSessionOptions
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

type N404ErrorCode string
type N404StatusCode *int
type N409ErrorCode string
type N409StatusCode *int

// Status returns HTTPResponse.Status
func (r RetrieveACheckoutSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveACheckoutSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAllCouponsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]Coupon `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllCouponsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllCouponsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateACouponResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Coupon
	JSON400      *interface{}
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateACouponResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateACouponResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveACouponResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Coupon
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r RetrieveACouponResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveACouponResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateACouponResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Coupon
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateACouponResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateACouponResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAllOrdersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]OrderOptions `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllOrdersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllOrdersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveAnOrderResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OrderOptions
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r RetrieveAnOrderResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveAnOrderResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelAnOrderResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Order
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
	JSON409 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N409ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N409StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r CancelAnOrderResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelAnOrderResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAllPaymentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]PaymentOptions `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllPaymentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllPaymentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveAPaymentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PaymentOptions
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r RetrieveAPaymentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveAPaymentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAPaymentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Payment
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
	JSON409 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N409ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N409StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateAPaymentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAPaymentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAllPromotionCodesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]PromotionCode `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllPromotionCodesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllPromotionCodesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAPromotionCodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PromotionCode
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
	JSON409 *struct {
		Details    *[]interface{} `json:"details,omitempty"`
		ErrorCode  *string        `json:"errorCode,omitempty"`
		Message    *string        `json:"message,omitempty"`
		StatusCode *int           `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateAPromotionCodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAPromotionCodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveAPromotionCodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PromotionCode
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r RetrieveAPromotionCodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveAPromotionCodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateAPromotionCodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PromotionCode
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateAPromotionCodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateAPromotionCodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAllRefundsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]RefundOptions `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ListAllRefundsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAllRefundsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateARefundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Refund
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON403 *struct {
		Details    *[]interface{} `json:"details,omitempty"`
		ErrorCode  *string        `json:"errorCode,omitempty"`
		Message    *string        `json:"message,omitempty"`
		StatusCode *int           `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []map[string]interface{} `json:"details"`
		ErrorCode  N404ErrorCode            `json:"errorCode"`
		Message    string                   `json:"message"`
		StatusCode N404StatusCode           `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateARefundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateARefundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RetrieveARefundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RefundOptions
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r RetrieveARefundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RetrieveARefundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateARefundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Refund
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{}  `json:"details"`
		ErrorCode  N404ErrorCode  `json:"errorCode"`
		Message    string         `json:"message"`
		StatusCode N404StatusCode `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateARefundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateARefundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetV1WebhookEndpointsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Data *[]WebhookEndpoint `json:"data,omitempty"`

		// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
		MaxResults *MaxResults `json:"maxResults,omitempty"`

		// The token for the next page of the collection of objects.
		NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

		// A string representing the object’s type. The value is always `collection` for collection objects.
		Object *CollectionObject `json:"object,omitempty"`

		// The token for the page of the collection of objects.
		PageToken *PageToken `json:"pageToken,omitempty"`

		// The actual number of objects returned for this call. This value is less than or equal to maxResults.
		Results *Results `json:"results,omitempty"`
	}
	JSON401 *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetV1WebhookEndpointsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetV1WebhookEndpointsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostV1WebhookEndpointsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WebhookEndpoint
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PostV1WebhookEndpointsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostV1WebhookEndpointsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetV1WebhookEndpointsWebhookEndpointIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WebhookEndpoint
	JSON401      *struct {
		Realm      *string `json:"realm,omitempty"`
		Scheme     *string `json:"scheme,omitempty"`
		StatusCode *int    `json:"statusCode,omitempty"`
	}
	JSON404 *struct {
		Details    []interface{} `json:"details"`
		ErrorCode  string        `json:"errorCode"`
		Message    string        `json:"message"`
		StatusCode float32       `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r GetV1WebhookEndpointsWebhookEndpointIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetV1WebhookEndpointsWebhookEndpointIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchV1WebhookEndpointsWebhookEndpointIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WebhookEndpoint
	JSON404      *struct {
		Details    []interface{} `json:"details"`
		ErrorCode  string        `json:"errorCode"`
		Message    string        `json:"message"`
		StatusCode float32       `json:"statusCode"`
	}
}

// Status returns HTTPResponse.Status
func (r PatchV1WebhookEndpointsWebhookEndpointIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchV1WebhookEndpointsWebhookEndpointIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListAllCheckoutSessionsWithResponse request returning *ListAllCheckoutSessionsResponse
func (c *ClientWithResponses) ListAllCheckoutSessionsWithResponse(ctx context.Context, params *ListAllCheckoutSessionsParams, reqEditors ...RequestEditorFn) (*ListAllCheckoutSessionsResponse, error) {
	rsp, err := c.ListAllCheckoutSessions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAllCheckoutSessionsResponse(rsp)
}

// CreateACheckoutSessionWithBodyWithResponse request with arbitrary body returning *CreateACheckoutSessionResponse
func (c *ClientWithResponses) CreateACheckoutSessionWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error) {
	rsp, err := c.CreateACheckoutSessionWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateACheckoutSessionResponse(rsp)
}

func (c *ClientWithResponses) CreateACheckoutSessionWithResponse(ctx context.Context, body CreateACheckoutSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACheckoutSessionResponse, error) {
	rsp, err := c.CreateACheckoutSession(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return ParseCreateACheckoutSessionResponse(rsp)
}

// RetrieveACheckoutSessionWithResponse request returning *RetrieveACheckoutSessionResponse
func (c *ClientWithResponses) RetrieveACheckoutSessionWithResponse(ctx context.Context, checkoutSessionId string, params *RetrieveACheckoutSessionParams, reqEditors ...RequestEditorFn) (*RetrieveACheckoutSessionResponse, error) {
	rsp, err := c.RetrieveACheckoutSession(ctx, checkoutSessionId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveACheckoutSessionResponse(rsp)
}

// ListAllCouponsWithResponse request returning *ListAllCouponsResponse
func (c *ClientWithResponses) ListAllCouponsWithResponse(ctx context.Context, params *ListAllCouponsParams, reqEditors ...RequestEditorFn) (*ListAllCouponsResponse, error) {
	rsp, err := c.ListAllCoupons(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}

	return ParseListAllCouponsResponse(rsp)
}

// CreateACouponWithBodyWithResponse request with arbitrary body returning *CreateACouponResponse
func (c *ClientWithResponses) CreateACouponWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error) {
	rsp, err := c.CreateACouponWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateACouponResponse(rsp)
}

func (c *ClientWithResponses) CreateACouponWithResponse(ctx context.Context, body CreateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateACouponResponse, error) {
	rsp, err := c.CreateACoupon(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateACouponResponse(rsp)
}

// RetrieveACouponWithResponse request returning *RetrieveACouponResponse
func (c *ClientWithResponses) RetrieveACouponWithResponse(ctx context.Context, couponId string, reqEditors ...RequestEditorFn) (*RetrieveACouponResponse, error) {
	rsp, err := c.RetrieveACoupon(ctx, couponId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveACouponResponse(rsp)
}

// UpdateACouponWithBodyWithResponse request with arbitrary body returning *UpdateACouponResponse
func (c *ClientWithResponses) UpdateACouponWithBodyWithResponse(ctx context.Context, couponId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error) {
	rsp, err := c.UpdateACouponWithBody(ctx, couponId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateACouponResponse(rsp)
}

func (c *ClientWithResponses) UpdateACouponWithResponse(ctx context.Context, couponId string, body UpdateACouponJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateACouponResponse, error) {
	rsp, err := c.UpdateACoupon(ctx, couponId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateACouponResponse(rsp)
}

// ListAllOrdersWithResponse request returning *ListAllOrdersResponse
func (c *ClientWithResponses) ListAllOrdersWithResponse(ctx context.Context, params *ListAllOrdersParams, reqEditors ...RequestEditorFn) (*ListAllOrdersResponse, error) {
	rsp, err := c.ListAllOrders(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAllOrdersResponse(rsp)
}

// RetrieveAnOrderWithResponse request returning *RetrieveAnOrderResponse
func (c *ClientWithResponses) RetrieveAnOrderWithResponse(ctx context.Context, orderId string, params *RetrieveAnOrderParams, reqEditors ...RequestEditorFn) (*RetrieveAnOrderResponse, error) {
	rsp, err := c.RetrieveAnOrder(ctx, orderId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveAnOrderResponse(rsp)
}

// CancelAnOrderWithResponse request returning *CancelAnOrderResponse
func (c *ClientWithResponses) CancelAnOrderWithResponse(ctx context.Context, orderId string, reqEditors ...RequestEditorFn) (*CancelAnOrderResponse, error) {
	rsp, err := c.CancelAnOrder(ctx, orderId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCancelAnOrderResponse(rsp)
}

// ListAllPaymentsWithResponse request returning *ListAllPaymentsResponse
func (c *ClientWithResponses) ListAllPaymentsWithResponse(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*ListAllPaymentsResponse, error) {
	rsp, err := c.ListAllPayments(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAllPaymentsResponse(rsp)
}

// RetrieveAPaymentWithResponse request returning *RetrieveAPaymentResponse
func (c *ClientWithResponses) RetrieveAPaymentWithResponse(ctx context.Context, paymentId string, params *RetrieveAPaymentParams, reqEditors ...RequestEditorFn) (*RetrieveAPaymentResponse, error) {
	rsp, err := c.RetrieveAPayment(ctx, paymentId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveAPaymentResponse(rsp)
}

// CreateAPaymentWithBodyWithResponse request with arbitrary body returning *CreateAPaymentResponse
func (c *ClientWithResponses) CreateAPaymentWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error) {
	rsp, err := c.CreateAPaymentWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAPaymentResponse(rsp)
}

func (c *ClientWithResponses) CreateAPaymentWithResponse(ctx context.Context, body CreateAPaymentJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPaymentResponse, error) {
	rsp, err := c.CreateAPayment(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAPaymentResponse(rsp)
}

// ListAllPromotionCodesWithResponse request returning *ListAllPromotionCodesResponse
func (c *ClientWithResponses) ListAllPromotionCodesWithResponse(ctx context.Context, params *ListAllPromotionCodesParams, reqEditors ...RequestEditorFn) (*ListAllPromotionCodesResponse, error) {
	rsp, err := c.ListAllPromotionCodes(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAllPromotionCodesResponse(rsp)
}

// CreateAPromotionCodeWithBodyWithResponse request with arbitrary body returning *CreateAPromotionCodeResponse
func (c *ClientWithResponses) CreateAPromotionCodeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error) {
	rsp, err := c.CreateAPromotionCodeWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAPromotionCodeResponse(rsp)
}

func (c *ClientWithResponses) CreateAPromotionCodeWithResponse(ctx context.Context, body CreateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAPromotionCodeResponse, error) {
	rsp, err := c.CreateAPromotionCode(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAPromotionCodeResponse(rsp)
}

// RetrieveAPromotionCodeWithResponse request returning *RetrieveAPromotionCodeResponse
func (c *ClientWithResponses) RetrieveAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, reqEditors ...RequestEditorFn) (*RetrieveAPromotionCodeResponse, error) {
	rsp, err := c.RetrieveAPromotionCode(ctx, promotionCodeId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveAPromotionCodeResponse(rsp)
}

// UpdateAPromotionCodeWithBodyWithResponse request with arbitrary body returning *UpdateAPromotionCodeResponse
func (c *ClientWithResponses) UpdateAPromotionCodeWithBodyWithResponse(ctx context.Context, promotionCodeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error) {
	rsp, err := c.UpdateAPromotionCodeWithBody(ctx, promotionCodeId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAPromotionCodeResponse(rsp)
}

func (c *ClientWithResponses) UpdateAPromotionCodeWithResponse(ctx context.Context, promotionCodeId string, body UpdateAPromotionCodeJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAPromotionCodeResponse, error) {
	rsp, err := c.UpdateAPromotionCode(ctx, promotionCodeId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAPromotionCodeResponse(rsp)
}

// ListAllRefundsWithResponse request returning *ListAllRefundsResponse
func (c *ClientWithResponses) ListAllRefundsWithResponse(ctx context.Context, params *ListAllRefundsParams, reqEditors ...RequestEditorFn) (*ListAllRefundsResponse, error) {
	rsp, err := c.ListAllRefunds(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAllRefundsResponse(rsp)
}

// CreateARefundWithBodyWithResponse request with arbitrary body returning *CreateARefundResponse
func (c *ClientWithResponses) CreateARefundWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error) {
	rsp, err := c.CreateARefundWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateARefundResponse(rsp)
}

func (c *ClientWithResponses) CreateARefundWithResponse(ctx context.Context, body CreateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateARefundResponse, error) {
	rsp, err := c.CreateARefund(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateARefundResponse(rsp)
}

// RetrieveARefundWithResponse request returning *RetrieveARefundResponse
func (c *ClientWithResponses) RetrieveARefundWithResponse(ctx context.Context, refundId string, params *RetrieveARefundParams, reqEditors ...RequestEditorFn) (*RetrieveARefundResponse, error) {
	rsp, err := c.RetrieveARefund(ctx, refundId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRetrieveARefundResponse(rsp)
}

// UpdateARefundWithBodyWithResponse request with arbitrary body returning *UpdateARefundResponse
func (c *ClientWithResponses) UpdateARefundWithBodyWithResponse(ctx context.Context, refundId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error) {
	rsp, err := c.UpdateARefundWithBody(ctx, refundId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateARefundResponse(rsp)
}

func (c *ClientWithResponses) UpdateARefundWithResponse(ctx context.Context, refundId string, body UpdateARefundJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateARefundResponse, error) {
	rsp, err := c.UpdateARefund(ctx, refundId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateARefundResponse(rsp)
}

// GetV1WebhookEndpointsWithResponse request returning *GetV1WebhookEndpointsResponse
func (c *ClientWithResponses) GetV1WebhookEndpointsWithResponse(ctx context.Context, params *GetV1WebhookEndpointsParams, reqEditors ...RequestEditorFn) (*GetV1WebhookEndpointsResponse, error) {
	rsp, err := c.GetV1WebhookEndpoints(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetV1WebhookEndpointsResponse(rsp)
}

// PostV1WebhookEndpointsWithBodyWithResponse request with arbitrary body returning *PostV1WebhookEndpointsResponse
func (c *ClientWithResponses) PostV1WebhookEndpointsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1WebhookEndpointsResponse, error) {
	rsp, err := c.PostV1WebhookEndpointsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1WebhookEndpointsResponse(rsp)
}

func (c *ClientWithResponses) PostV1WebhookEndpointsWithResponse(ctx context.Context, body PostV1WebhookEndpointsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1WebhookEndpointsResponse, error) {
	rsp, err := c.PostV1WebhookEndpoints(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostV1WebhookEndpointsResponse(rsp)
}

// GetV1WebhookEndpointsWebhookEndpointIdWithResponse request returning *GetV1WebhookEndpointsWebhookEndpointIdResponse
func (c *ClientWithResponses) GetV1WebhookEndpointsWebhookEndpointIdWithResponse(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*GetV1WebhookEndpointsWebhookEndpointIdResponse, error) {
	rsp, err := c.GetV1WebhookEndpointsWebhookEndpointId(ctx, webhookEndpointId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetV1WebhookEndpointsWebhookEndpointIdResponse(rsp)
}

// PatchV1WebhookEndpointsWebhookEndpointIdWithBodyWithResponse request with arbitrary body returning *PatchV1WebhookEndpointsWebhookEndpointIdResponse
func (c *ClientWithResponses) PatchV1WebhookEndpointsWebhookEndpointIdWithBodyWithResponse(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchV1WebhookEndpointsWebhookEndpointIdResponse, error) {
	rsp, err := c.PatchV1WebhookEndpointsWebhookEndpointIdWithBody(ctx, webhookEndpointId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchV1WebhookEndpointsWebhookEndpointIdResponse(rsp)
}

func (c *ClientWithResponses) PatchV1WebhookEndpointsWebhookEndpointIdWithResponse(ctx context.Context, webhookEndpointId string, body PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchV1WebhookEndpointsWebhookEndpointIdResponse, error) {
	rsp, err := c.PatchV1WebhookEndpointsWebhookEndpointId(ctx, webhookEndpointId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchV1WebhookEndpointsWebhookEndpointIdResponse(rsp)
}

// ParseListAllCheckoutSessionsResponse parses an HTTP response from a ListAllCheckoutSessionsWithResponse call
func ParseListAllCheckoutSessionsResponse(rsp *http.Response) (*ListAllCheckoutSessionsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllCheckoutSessionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]CheckoutSessionOptions `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCreateACheckoutSessionResponse parses an HTTP response from a CreateACheckoutSessionWithResponse call
func ParseCreateACheckoutSessionResponse(rsp *http.Response) (*CreateACheckoutSessionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)

	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateACheckoutSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CheckoutSessionExpanded
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseRetrieveACheckoutSessionResponse parses an HTTP response from a RetrieveACheckoutSessionWithResponse call
func ParseRetrieveACheckoutSessionResponse(rsp *http.Response) (*RetrieveACheckoutSessionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveACheckoutSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CheckoutSessionOptions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseListAllCouponsResponse parses an HTTP response from a ListAllCouponsWithResponse call
func ParseListAllCouponsResponse(rsp *http.Response) (*ListAllCouponsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllCouponsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]Coupon `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCreateACouponResponse parses an HTTP response from a CreateACouponWithResponse call
func ParseCreateACouponResponse(rsp *http.Response) (*CreateACouponResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateACouponResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Coupon
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseRetrieveACouponResponse parses an HTTP response from a RetrieveACouponWithResponse call
func ParseRetrieveACouponResponse(rsp *http.Response) (*RetrieveACouponResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveACouponResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Coupon
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateACouponResponse parses an HTTP response from a UpdateACouponWithResponse call
func ParseUpdateACouponResponse(rsp *http.Response) (*UpdateACouponResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateACouponResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Coupon
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseListAllOrdersResponse parses an HTTP response from a ListAllOrdersWithResponse call
func ParseListAllOrdersResponse(rsp *http.Response) (*ListAllOrdersResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllOrdersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]OrderOptions `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseRetrieveAnOrderResponse parses an HTTP response from a RetrieveAnOrderWithResponse call
func ParseRetrieveAnOrderResponse(rsp *http.Response) (*RetrieveAnOrderResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveAnOrderResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OrderOptions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCancelAnOrderResponse parses an HTTP response from a CancelAnOrderWithResponse call
func ParseCancelAnOrderResponse(rsp *http.Response) (*CancelAnOrderResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CancelAnOrderResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Order
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N409ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N409StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseListAllPaymentsResponse parses an HTTP response from a ListAllPaymentsWithResponse call
func ParseListAllPaymentsResponse(rsp *http.Response) (*ListAllPaymentsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllPaymentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]PaymentOptions `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseRetrieveAPaymentResponse parses an HTTP response from a RetrieveAPaymentWithResponse call
func ParseRetrieveAPaymentResponse(rsp *http.Response) (*RetrieveAPaymentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveAPaymentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PaymentOptions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseCreateAPaymentResponse parses an HTTP response from a CreateAPaymentWithResponse call
func ParseCreateAPaymentResponse(rsp *http.Response) (*CreateAPaymentResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAPaymentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Payment
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N409ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N409StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseListAllPromotionCodesResponse parses an HTTP response from a ListAllPromotionCodesWithResponse call
func ParseListAllPromotionCodesResponse(rsp *http.Response) (*ListAllPromotionCodesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllPromotionCodesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]PromotionCode `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCreateAPromotionCodeResponse parses an HTTP response from a CreateAPromotionCodeWithResponse call
func ParseCreateAPromotionCodeResponse(rsp *http.Response) (*CreateAPromotionCodeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAPromotionCodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PromotionCode
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest struct {
			Details    *[]interface{} `json:"details,omitempty"`
			ErrorCode  *string        `json:"errorCode,omitempty"`
			Message    *string        `json:"message,omitempty"`
			StatusCode *int           `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ParseRetrieveAPromotionCodeResponse parses an HTTP response from a RetrieveAPromotionCodeWithResponse call
func ParseRetrieveAPromotionCodeResponse(rsp *http.Response) (*RetrieveAPromotionCodeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveAPromotionCodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PromotionCode
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateAPromotionCodeResponse parses an HTTP response from a UpdateAPromotionCodeWithResponse call
func ParseUpdateAPromotionCodeResponse(rsp *http.Response) (*UpdateAPromotionCodeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateAPromotionCodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PromotionCode
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseListAllRefundsResponse parses an HTTP response from a ListAllRefundsWithResponse call
func ParseListAllRefundsResponse(rsp *http.Response) (*ListAllRefundsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAllRefundsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]RefundOptions `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCreateARefundResponse parses an HTTP response from a CreateARefundWithResponse call
func ParseCreateARefundResponse(rsp *http.Response) (*CreateARefundResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateARefundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Refund
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			Details    *[]interface{} `json:"details,omitempty"`
			ErrorCode  *string        `json:"errorCode,omitempty"`
			Message    *string        `json:"message,omitempty"`
			StatusCode *int           `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []map[string]interface{} `json:"details"`
			ErrorCode  N404ErrorCode            `json:"errorCode"`
			Message    string                   `json:"message"`
			StatusCode N404StatusCode           `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseRetrieveARefundResponse parses an HTTP response from a RetrieveARefundWithResponse call
func ParseRetrieveARefundResponse(rsp *http.Response) (*RetrieveARefundResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RetrieveARefundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RefundOptions
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateARefundResponse parses an HTTP response from a UpdateARefundWithResponse call
func ParseUpdateARefundResponse(rsp *http.Response) (*UpdateARefundResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateARefundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Refund
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{}  `json:"details"`
			ErrorCode  N404ErrorCode  `json:"errorCode"`
			Message    string         `json:"message"`
			StatusCode N404StatusCode `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetV1WebhookEndpointsResponse parses an HTTP response from a GetV1WebhookEndpointsWithResponse call
func ParseGetV1WebhookEndpointsResponse(rsp *http.Response) (*GetV1WebhookEndpointsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetV1WebhookEndpointsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Data *[]WebhookEndpoint `json:"data,omitempty"`

			// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
			MaxResults *MaxResults `json:"maxResults,omitempty"`

			// The token for the next page of the collection of objects.
			NextPageToken *NextPageToken `json:"nextPageToken,omitempty"`

			// A string representing the object’s type. The value is always `collection` for collection objects.
			Object *CollectionObject `json:"object,omitempty"`

			// The token for the page of the collection of objects.
			PageToken *PageToken `json:"pageToken,omitempty"`

			// The actual number of objects returned for this call. This value is less than or equal to maxResults.
			Results *Results `json:"results,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParsePostV1WebhookEndpointsResponse parses an HTTP response from a PostV1WebhookEndpointsWithResponse call
func ParsePostV1WebhookEndpointsResponse(rsp *http.Response) (*PostV1WebhookEndpointsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostV1WebhookEndpointsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WebhookEndpoint
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseGetV1WebhookEndpointsWebhookEndpointIdResponse parses an HTTP response from a GetV1WebhookEndpointsWebhookEndpointIdWithResponse call
func ParseGetV1WebhookEndpointsWebhookEndpointIdResponse(rsp *http.Response) (*GetV1WebhookEndpointsWebhookEndpointIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetV1WebhookEndpointsWebhookEndpointIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WebhookEndpoint
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest struct {
			Realm      *string `json:"realm,omitempty"`
			Scheme     *string `json:"scheme,omitempty"`
			StatusCode *int    `json:"statusCode,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{} `json:"details"`
			ErrorCode  string        `json:"errorCode"`
			Message    string        `json:"message"`
			StatusCode float32       `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParsePatchV1WebhookEndpointsWebhookEndpointIdResponse parses an HTTP response from a PatchV1WebhookEndpointsWebhookEndpointIdWithResponse call
func ParsePatchV1WebhookEndpointsWebhookEndpointIdResponse(rsp *http.Response) (*PatchV1WebhookEndpointsWebhookEndpointIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchV1WebhookEndpointsWebhookEndpointIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WebhookEndpoint
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest struct {
			Details    []interface{} `json:"details"`
			ErrorCode  string        `json:"errorCode"`
			Message    string        `json:"message"`
			StatusCode float32       `json:"statusCode"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}
