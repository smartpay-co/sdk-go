package smartpay

const (
	ApiEndpoint = "https://api.smartpay.co"
)

// Defines values for AddressCountry.
const (
	AddressCountryJP AddressCountry = "JP"
)

// Defines values for AddressType.
const (
	AddressTypeGift AddressType = "gift"

	AddressTypeHome AddressType = "home"

	AddressTypeLocker AddressType = "locker"

	AddressTypeOffice AddressType = "office"

	AddressTypeStore AddressType = "store"
)

// Defines values for CaptureMethod.
const (
	CaptureMethodAutomatic CaptureMethod = "automatic"

	CaptureMethodManual CaptureMethod = "manual"
)

// Defines values for CollectionObject.
const (
	CollectionObjectCollection CollectionObject = "collection"
)

// Defines values for Currency.
const (
	CurrencyJPY Currency = "JPY"
)

// Defines values for CustomerInfoLegalGender.
const (
	CustomerInfoLegalGenderFemale CustomerInfoLegalGender = "female"

	CustomerInfoLegalGenderMale CustomerInfoLegalGender = "male"
)

// Defines values for DiscountType.
const (
	DiscountTypeAmount DiscountType = "amount"

	DiscountTypePercentage DiscountType = "percentage"
)

// Defines values for EventSubscription.
const (
	EventSubscriptionCouponCreated EventSubscription = "coupon.created"

	EventSubscriptionCouponUpdated EventSubscription = "coupon.updated"

	EventSubscriptionMerchantUserCreated EventSubscription = "merchant_user.created"

	EventSubscriptionMerchantUserPasswordReset EventSubscription = "merchant_user.password_reset"

	EventSubscriptionOrderAuthorized EventSubscription = "order.authorized"

	EventSubscriptionOrderCanceled EventSubscription = "order.canceled"

	EventSubscriptionOrderRejected EventSubscription = "order.rejected"

	EventSubscriptionPaymentCreated EventSubscription = "payment.created"

	EventSubscriptionPayoutGenerated EventSubscription = "payout.generated"

	EventSubscriptionPayoutPaid EventSubscription = "payout.paid"

	EventSubscriptionPromotionCodeCreated EventSubscription = "promotion_code.created"

	EventSubscriptionPromotionCodeUpdated EventSubscription = "promotion_code.updated"

	EventSubscriptionRefundCreated EventSubscription = "refund.created"

	EventSubscriptionTokenApproved EventSubscription = "token.approved"

	EventSubscriptionTokenRejected EventSubscription = "token.rejected"

	EventSubscriptionTokenDeleted EventSubscription = "token.deleted"

	EventSubscriptionTokenDisabled EventSubscription = "token.disabled"

	EventSubscriptionTokenEnabled EventSubscription = "token.enabled"

	EventSubscriptionTokenCreated EventSubscription = "token.created"

	EventSubscriptionTokenUsed EventSubscription = "token.used"
)

// Defines values for Locale.
const (
	LocaleEn Locale = "en"

	LocaleJa Locale = "ja"
)

// Defines values for OrderStatus.
const (
	OrderStatusCanceled OrderStatus = "canceled"

	OrderStatusPending OrderStatus = "pending"

	OrderStatusRejected OrderStatus = "rejected"

	OrderStatusRequiresAuthorization OrderStatus = "requires_authorization"

	OrderStatusRequiresCapture OrderStatus = "requires_capture"

	OrderStatusSucceeded OrderStatus = "succeeded"
)

// Defines values for PaymentStatus.
const (
	PaymentStatusProcessed PaymentStatus = "processed"

	PaymentStatusRefunded PaymentStatus = "refunded"
)

// Defines values for PaymentExpandedStatus.
const (
	PaymentExpandedStatusProcessed PaymentExpandedStatus = "processed"

	PaymentExpandedStatusRefunded PaymentExpandedStatus = "refunded"
)

// Defines values for PaymentCreateCancelRemainder
const (
	PaymentCreateCancelRemainderAutomatic PaymentCreateCancelRemainder = "automatic"
	PaymentCreateCancelRemainderManual    PaymentCreateCancelRemainder = "manual"
)

// Defines values for RefundReason.
const (
	RefundReasonFraudulent RefundReason = "fraudulent"

	RefundReasonRequestedByCustomer RefundReason = "requested_by_customer"
)

// Defines values for RefundStatus.
const (
	RefundStatusFailed RefundStatus = "failed"

	RefundStatusSucceeded RefundStatus = "succeeded"
)

// Defines values for RefundExpandedReason.
const (
	RefundExpandedReasonFraudulent RefundExpandedReason = "fraudulent"

	RefundExpandedReasonRequestedByCustomer RefundExpandedReason = "requested_by_customer"
)

// Defines values for RefundExpandedStatus.
const (
	RefundExpandedStatusFailed RefundExpandedStatus = "failed"

	RefundExpandedStatusSucceeded RefundExpandedStatus = "succeeded"
)

// Defines values for Expand.
const (
	ExpandAll Expand = "all"

	ExpandNo Expand = "no"
)

// Defines values for LineItemKind.
const (
	LineItemKindProduct  LineItemKind = "product"
	LineItemKindDiscount LineItemKind = "discount"
	LineItemKindTax      LineItemKind = "tax"
)

// Defines values for Mode.
const (
	ModeToken Mode = "token"
)

// Defines values for TokenStatus.
const (
	TokenStatusRequiresAuthorization TokenStatus = "requires_authorization"
	TokenStatusActive                TokenStatus = "active"
	TokenStatusRejected              TokenStatus = "rejected"
	TokenStatusDisabled              TokenStatus = "disabled"
	TokenStatusDeleted               TokenStatus = "deleted"
)
