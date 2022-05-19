package smartpay

import (
	"context"
	"io"
	"net/http"
)

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

	// ListAllPayments request
	ListAllPayments(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

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

	// GetV1WebhookEndpoints request
	GetV1WebhookEndpoints(ctx context.Context, params *GetV1WebhookEndpointsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostV1WebhookEndpoints request with any body
	PostV1WebhookEndpointsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostV1WebhookEndpoints(ctx context.Context, body PostV1WebhookEndpointsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetV1WebhookEndpointsWebhookEndpointId request
	GetV1WebhookEndpointsWebhookEndpointId(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchV1WebhookEndpointsWebhookEndpointId request with any body
	PatchV1WebhookEndpointsWebhookEndpointIdWithBody(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchV1WebhookEndpointsWebhookEndpointId(ctx context.Context, webhookEndpointId string, body PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
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

	// ListAllPayments request
	ListAllPaymentsWithResponse(ctx context.Context, params *ListAllPaymentsParams, reqEditors ...RequestEditorFn) (*ListAllPaymentsResponse, error)

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

	// GetV1WebhookEndpoints request
	GetV1WebhookEndpointsWithResponse(ctx context.Context, params *GetV1WebhookEndpointsParams, reqEditors ...RequestEditorFn) (*GetV1WebhookEndpointsResponse, error)

	// PostV1WebhookEndpoints request with any body
	PostV1WebhookEndpointsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostV1WebhookEndpointsResponse, error)

	PostV1WebhookEndpointsWithResponse(ctx context.Context, body PostV1WebhookEndpointsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostV1WebhookEndpointsResponse, error)

	// GetV1WebhookEndpointsWebhookEndpointId request
	GetV1WebhookEndpointsWebhookEndpointIdWithResponse(ctx context.Context, webhookEndpointId string, reqEditors ...RequestEditorFn) (*GetV1WebhookEndpointsWebhookEndpointIdResponse, error)

	// PatchV1WebhookEndpointsWebhookEndpointId request with any body
	PatchV1WebhookEndpointsWebhookEndpointIdWithBodyWithResponse(ctx context.Context, webhookEndpointId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchV1WebhookEndpointsWebhookEndpointIdResponse, error)

	PatchV1WebhookEndpointsWebhookEndpointIdWithResponse(ctx context.Context, webhookEndpointId string, body PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchV1WebhookEndpointsWebhookEndpointIdResponse, error)
}
