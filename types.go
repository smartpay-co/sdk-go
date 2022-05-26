package smartpay

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Address
type Address struct {
	// The province, state, prefecture, county, etc. of a location. In Japan it refers to the prefecture (e.g. 東京都).
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// The country as represented by the two-letter ISO 3166-1 code.
	Country AddressCountry `json:"country"`

	// Street address.
	Line1 string `json:"line1"`

	// Building name and room number.
	Line2 *string `json:"line2,omitempty"`
	Line3 *string `json:"line3,omitempty"`
	Line4 *string `json:"line4,omitempty"`
	Line5 *string `json:"line5,omitempty"`

	// The city or town of a location, with a maximum of 80 characters (e.g. 目黒区).
	Locality string `json:"locality"`

	// The Postal Code.
	PostalCode  string  `json:"postalCode"`
	SubLocality *string `json:"subLocality,omitempty"`
}

// The country as represented by the two-letter ISO 3166-1 code.
type AddressCountry string

// Address Type
type AddressType string

// The URL the customer will be redirected to if the Checkout Session hasn't completed successfully. This means the Checkout failed, or the customer decided to cancel it.
type CancelUrl string

// CaptureMethod defines model for captureMethod.
type CaptureMethod string

// The URL to launch Smartpay checkout for this checkout session. Redirect your customer to this URL to complete the purchase.
type CheckoutSessioUrl string

// Checkout Session
type CheckoutSession struct {
	// The URL the customer will be redirected to if the Checkout Session hasn't completed successfully. This means the Checkout failed, or the customer decided to cancel it.
	CancelUrl *CancelUrl `json:"cancelUrl,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Customer Information, the details provided here are used to pre-populate forms for your customer's checkout experiences. The more information you send, the better experience the customer will have.
	CustomerInfo *CustomerInfo `json:"customerInfo,omitempty"`

	// The unique identifier for the Checkout Session object.
	Id *CheckoutSessionId `json:"id,omitempty"`

	// Locale
	Locale *Locale `json:"locale,omitempty"`

	// A string representing the object’s type. The value is always `checkoutSession` for Checkout Session objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Order *OrderId `json:"order,omitempty"`

	// The URL the customer will be redirected to if the Checkout Session completed successfully. This means the Checkout succeeded, i.e. the customer authorized the order.
	SuccessUrl *SuccessUrl `json:"successUrl,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`

	// The URL to launch Smartpay checkout for this checkout session. Redirect your customer to this URL to complete the purchase.
	Url *CheckoutSessioUrl `json:"url,omitempty"`
}

// Expanded Checkout Session
type CheckoutSessionExpanded struct {
	// The URL the customer will be redirected to if the Checkout Session hasn't completed successfully. This means the Checkout failed, or the customer decided to cancel it.
	CancelUrl *CancelUrl `json:"cancelUrl,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Customer Information, the details provided here are used to pre-populate forms for your customer's checkout experiences. The more information you send, the better experience the customer will have.
	CustomerInfo *CustomerInfo `json:"customerInfo,omitempty"`

	// The unique identifier for the Checkout Session object.
	Id *CheckoutSessionId `json:"id,omitempty"`

	// Locale
	Locale *Locale `json:"locale,omitempty"`

	// A string representing the object’s type. The value is always `checkoutSession` for Checkout Session objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Order *OrderExpanded `json:"order,omitempty"`

	// The URL the customer will be redirected to if the Checkout Session completed successfully. This means the Checkout succeeded, i.e. the customer authorized the order.
	SuccessUrl *SuccessUrl `json:"successUrl,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`

	// The URL to launch Smartpay checkout for this checkout session. Redirect your customer to this URL to complete the purchase.
	Url *CheckoutSessioUrl `json:"url,omitempty"`
}

// The unique identifier for the Checkout Session object.
type CheckoutSessionId string

// CheckoutSessionOptions defines model for checkoutSessionOptions.
type CheckoutSessionOptions interface{}

// A string representing the object’s type. The value is always `collection` for collection objects.
type CollectionObject string

// Coupon
type Coupon struct {
	// Has the value `true` if the coupon is active and can be used, or the value `false` if it is not.
	Active *bool `json:"active,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// If the `discountType` is `amount`, the discount amount of this coupon object.
	DiscountAmount *float32 `json:"discountAmount,omitempty"`

	// If the `discountType` is `percentage`, the discount percentage of this coupon object.
	DiscountPercentage *float32 `json:"discountPercentage,omitempty"`

	// Discount Type
	DiscountType *DiscountType `json:"discountType,omitempty"`

	// Time at which the Coupon expires. Measured in seconds since the Unix epoch.
	ExpiresAt *int `json:"expiresAt,omitempty"`

	// The unique identifier for the Coupon object.
	Id *CouponId `json:"id,omitempty"`

	// Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptionCount *int `json:"maxRedemptionCount,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The coupon's name, meant to be displayable to the customer.
	Name *string `json:"name,omitempty"`

	// A string representing the object’s type. The value is always `coupon` for Coupon objects.
	Object *string `json:"object,omitempty"`

	// Number of times this coupon has been applied to an order.
	RedemptionCount *int `json:"redemptionCount,omitempty"`

	// Has the value `true` if the coupon is Smartpay funded, or the value `false` if it is not.
	Sponsored *bool `json:"sponsored,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// The unique identifier for the Coupon object.
type CouponId string

// Time at which the object was created. Measured in milliseconds since the Unix epoch.
type CreatedAt int

// Three-letter ISO currency code, in uppercase. Must be a supported currency.
type Currency string

// Customer Information, the details provided here are used to pre-populate forms for your customer's checkout experiences. The more information you send, the better experience the customer will have.
type CustomerInfo struct {
	// The age of the customer's account on your website in days.
	AccountAge *int `json:"accountAge,omitempty"`

	// Address
	Address *Address `json:"address,omitempty"`

	// The date of birth of your customer. The date of birth specified here will be used to pre-populate the date of birth field in your customer's checkout experiences.
	DateOfBirth *openapi_types.Date `json:"dateOfBirth,omitempty"`

	// The email address of your customer. The email address specified here will be used to pre-populate the email address field in your customer's checkout experiences.
	EmailAddress *openapi_types.Email `json:"emailAddress,omitempty"`

	// The first name of your customer. The name specified here will be used to pre-populate the first name field in your customer's checkout experiences.
	FirstName *string `json:"firstName,omitempty"`

	// The kana version of the first name of your customer. The name specified here will be used to pre-populate the kana version of the first name field in your customer's checkout experiences.
	FirstNameKana *string `json:"firstNameKana,omitempty"`

	// The last name of your customer. The name specified here will be used to pre-populate the last name field in your customer's checkout experiences.
	LastName *string `json:"lastName,omitempty"`

	// The kana version of the last name of your customer. The name specified here will be used to pre-populate the kana version of the last name field in your customer's checkout experiences.
	LastNameKana *string `json:"lastNameKana,omitempty"`

	// The (legal) gender of your customer. The gender specified here will be used to pre-populate the gender field in your customer's checkout experiences.
	LegalGender *CustomerInfoLegalGender `json:"legalGender,omitempty"`

	// The phone number of your customer. The phone number specified here will be used to pre-populate the phone number field in your customer's checkout experiences.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// The ID of the user in your system
	Reference *string `json:"reference,omitempty"`
}

// The (legal) gender of your customer. The gender specified here will be used to pre-populate the gender field in your customer's checkout experiences.
type CustomerInfoLegalGender string

// Discount
type Discount struct {
	Amount *float32 `json:"amount,omitempty"`

	// The unique identifier for the Coupon object.
	Coupon *CouponId `json:"coupon,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// The unique identifier for the Discount object.
	Id *DiscountId `json:"id,omitempty"`

	// A string representing the object’s type. The value is always `discount` for Discount objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Promotion Code object.
	PromotionCode *PromotionCodeId `json:"promotionCode,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`
}

// The discount amount applied to the order through a Smartpay coupon
//type DiscountAmount int
type DiscountAmount float32

// The unique identifier for the Discount object.
type DiscountId string

// Discount Type
type DiscountType string

// Even name to subscribe to
type EventSubscription string

// A URL for an image for this product, meant to be displayed to the customer.
type ImageUrl string

// Item
type Item struct {
	// The unit amount of this line item.
	Amount float32 `json:"amount"`

	// The brand of the Product.
	Brand *string `json:"brand,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency Currency `json:"currency"`

	// An arbitrary - ideally descriptive - long form explanation of the Line Item, meant to be displayed to the customer.
	Description *string `json:"description,omitempty"`

	// The Global Trade Item Number of the Product.
	Gtin *string `json:"gtin,omitempty"`

	// A list of up to 8 URLs of images for this product, meant to be displayed to the customer.
	Images *[]ImageUrl `json:"images,omitempty"`

	// A brief description of the price, not visible to customers.
	Label *string `json:"label,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The product’s name, meant to be displayed to the customer.
	Name             string  `json:"name"`
	PriceDescription *string `json:"priceDescription,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	PriceMetadata *Metadata `json:"priceMetadata,omitempty"`

	// An arbitrary - ideally descriptive - long form explanation of the Product, meant to be displayed to the customer.
	ProductDescription *string `json:"productDescription,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	ProductMetadata *Metadata `json:"productMetadata,omitempty"`

	// The quantity of products. Needs to be positive or zero.
	Quantity int `json:"quantity"`

	// A - ideally unique - string to reference the Product in your system (e.g. a product ID, etc.).
	Reference *string `json:"reference,omitempty"`

	// A URL of the publicly accessible page for this Product on your site or store.
	Url *ProductUrl `json:"url,omitempty"`
}

// Line Item
type LineItemExpanded struct {
	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt   *CreatedAt `json:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty"`

	// The unique identifier for the Line Item object.
	Id *LineItemId `json:"id,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `lineItem` for Line Item objects.
	Object *string `json:"object,omitempty"`

	// Price
	Price    *PriceExpanded `json:"price,omitempty"`
	Quantity *int           `json:"quantity,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// The unique identifier for the Line Item object.
type LineItemId string

// Locale
type Locale string

// The maximum number of objects returned for this call. Equals to the maxResults query parameter specified (or 20 if not specified).
type MaxResults int

// Set of up to 20 key-value pairs that you can attach to the object.
type Metadata map[string]interface{}

// The token for the next page of the collection of objects.
type NextPageToken string

// Order
type Order struct {
	// The amount intended to be collected through this order. A positive integer in the smallest currency unit.
	Amount *OrderAmount `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// The discount amount applied to the order through a Smartpay coupon
	DiscountAmount *DiscountAmount `json:"discountAmount,omitempty"`
	Discounts      *[]DiscountId   `json:"discounts,omitempty"`
	ExpiresAt      *int            `json:"expiresAt,omitempty"`

	// The unique identifier for the Payment object.
	Id        *OrderId      `json:"id,omitempty"`
	LineItems *[]LineItemId `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `order` for Order objects.
	Object    *string      `json:"object,omitempty"`
	Payments  *[]PaymentId `json:"payments,omitempty"`
	Reference *string      `json:"reference,omitempty"`

	// Shipping Information
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`

	// Order Status
	Status *OrderStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// The amount intended to be collected through this order. A positive integer in the smallest currency unit.
type OrderAmount float32

// Expanded Order
type OrderExpanded struct {
	// The amount intended to be collected through this order. A positive integer in the smallest currency unit.
	Amount *OrderAmount `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency    *Currency `json:"currency,omitempty"`
	Description *string   `json:"description,omitempty"`

	// The discount amount applied to the order through a Smartpay coupon
	DiscountAmount *DiscountAmount `json:"discountAmount,omitempty"`
	Discounts      *[]Discount     `json:"discounts,omitempty"`
	ExpiresAt      *int            `json:"expiresAt,omitempty"`

	// The unique identifier for the Payment object.
	Id        *OrderId            `json:"id,omitempty"`
	LineItems *[]LineItemExpanded `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `order` for Order objects.
	Object   *string            `json:"object,omitempty"`
	Payments *[]PaymentExpanded `json:"payments,omitempty"`

	// Shipping Information
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`

	// Order Status
	Status *OrderStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// The unique identifier for the Payment object.
type OrderId string

// Expandable Order
type OrderOptions interface{}

// Order Status
type OrderStatus string

// The token for the page of the collection of objects.
type PageToken string

// Payment
type Payment struct {
	Amount *float32 `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency    *Currency `json:"currency,omitempty"`
	Description *string   `json:"description,omitempty"`

	// The unique identifier for the Payment object.
	Id        *PaymentId    `json:"id,omitempty"`
	LineItems *[]LineItemId `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `payment` for Payment objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Order     *OrderId       `json:"order,omitempty"`
	Reference *string        `json:"reference,omitempty"`
	Status    *PaymentStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// PaymentStatus defines model for Payment.Status.
type PaymentStatus string

// Expanded Payment
type PaymentExpanded struct {
	Amount *float32 `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency    *Currency `json:"currency,omitempty"`
	Description *string   `json:"description,omitempty"`

	// The unique identifier for the Payment object.
	Id        *PaymentId          `json:"id,omitempty"`
	LineItems *[]LineItemExpanded `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `payment` for Payment objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Order     *OrderId               `json:"order,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Status    *PaymentExpandedStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// PaymentExpandedStatus defines model for PaymentExpanded.Status.
type PaymentExpandedStatus string

// The unique identifier for the Payment object.
type PaymentId string

// Expandable Payment
type PaymentOptions interface{}

// Price
type PriceExpanded struct {
	Active      *bool    `json:"active,omitempty"`
	Amount      *float32 `json:"amount,omitempty"`
	CreatedAt   *int     `json:"createdAt,omitempty"`
	Description *string  `json:"description,omitempty"`

	// The unique identifier for the Price object.
	Id    *PriceId `json:"id,omitempty"`
	Label *string  `json:"label,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `price` for Price objects.
	Object *string `json:"object,omitempty"`

	// Product
	Product *Product `json:"product,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test      *TestFlag `json:"test,omitempty"`
	UpdatedAt *int      `json:"updatedAt,omitempty"`
}

// The unique identifier for the Price object.
type PriceId string

// Product
type Product struct {
	Active     *bool              `json:"active,omitempty"`
	Brand      *string            `json:"brand,omitempty"`
	Categories *[]ProductCategory `json:"categories,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt   *CreatedAt `json:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty"`
	Gtin        *string    `json:"gtin,omitempty"`

	// The unique identifier for the Product object.
	Id     *ProductId  `json:"id,omitempty"`
	Images *[]ImageUrl `json:"images,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`
	Name     *string   `json:"name,omitempty"`

	// A string representing the object’s type. The value is always `product` for Product objects.
	Object    *string `json:"object,omitempty"`
	Reference *string `json:"reference,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`

	// A URL of the publicly accessible page for this Product on your site or store.
	Url *ProductUrl `json:"url,omitempty"`
}

// A category of items the product belongs to. Examples: apparel, cosmetics, etc.
type ProductCategory string

// The unique identifier for the Product object.
type ProductId string

// A URL of the publicly accessible page for this Product on your site or store.
type ProductUrl string

// A promotion code points to a coupon, and can be used to (attempt to) attach that coupon to an order.
type PromotionCode struct {
	// Flag indicating whether the promotion code is currently active.
	Active *bool `json:"active,omitempty"`

	// The customer-facing code. This code must be unique across all your promotion codes.
	Code *string `json:"code,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// The moment at which the promotion code can no longer be redeemed. Measured in seconds since the Unix epoch.
	ExpiresAt *int `json:"expiresAt,omitempty"`

	// A flag indicating that the Promotion Code should only be redeemed by Customer without any successful Smartpay payments with you.
	FirstTimeTransaction *bool `json:"firstTimeTransaction,omitempty"`

	// The unique identifier for the Promotion Code object.
	Id *PromotionCodeId `json:"id,omitempty"`

	// The maximum number of times this promotion code can be redeemedcan be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptionCount *int `json:"maxRedemptionCount,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// Minimum amount required to redeem this Promotion Code into a Coupon discount.
	MinimumAmount *float32 `json:"minimumAmount,omitempty"`

	// A string representing the object’s type. The value is always `promotionCode` for Promotion Code objects.
	Object *string `json:"object,omitempty"`

	// Flag indicating that the Promotion Code should only be redeemed once by a single Customer.
	OnePerCustomer *bool `json:"onePerCustomer,omitempty"`

	// The number of times this promotion code has been used.
	RedemptionCount *int `json:"redemptionCount,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// The unique identifier for the Promotion Code object.
type PromotionCodeId string

// Refund
type Refund struct {
	Amount *int `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency    *Currency `json:"currency,omitempty"`
	Description *string   `json:"description,omitempty"`

	// The unique identifier for the Refund object.
	Id        *RefundId     `json:"id,omitempty"`
	LineItems *[]LineItemId `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `refund` for Refund objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Payment   *PaymentId    `json:"payment,omitempty"`
	Reason    *RefundReason `json:"reason,omitempty"`
	Reference *string       `json:"reference,omitempty"`
	Status    *RefundStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// RefundReason defines model for Refund.Reason.
type RefundReason string

// RefundStatus defines model for Refund.Status.
type RefundStatus string

// Expanded Refund
type RefundExpanded struct {
	Amount *int `json:"amount,omitempty"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt *CreatedAt `json:"createdAt,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency    *Currency `json:"currency,omitempty"`
	Description *string   `json:"description,omitempty"`

	// The unique identifier for the Refund object.
	Id        *RefundId           `json:"id,omitempty"`
	LineItems *[]LineItemExpanded `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `refund` for Refund objects.
	Object *string `json:"object,omitempty"`

	// The unique identifier for the Payment object.
	Payment   *PaymentId            `json:"payment,omitempty"`
	Reason    *RefundExpandedReason `json:"reason,omitempty"`
	Reference *string               `json:"reference,omitempty"`
	Status    *RefundExpandedStatus `json:"status,omitempty"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test *TestFlag `json:"test,omitempty"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt *UpdatedAt `json:"updatedAt,omitempty"`
}

// RefundExpandedReason defines model for RefundExpanded.Reason.
type RefundExpandedReason string

// RefundExpandedStatus defines model for RefundExpanded.Status.
type RefundExpandedStatus string

// The unique identifier for the Refund object.
type RefundId string

// Expandable Refund
type RefundOptions interface{}

// The actual number of objects returned for this call. This value is less than or equal to maxResults.
type Results int

// Shipping Information
type ShippingInfo struct {
	// Address
	Address Address `json:"address"`

	// Address Type
	AddressType *AddressType `json:"addressType,omitempty"`

	// The delivery service that shipped a physical product, such as Yamato, Seino, Fedex, UPS, etc.
	CarrierName *string `json:"carrierName,omitempty"`

	// The shipping fee.
	FeeAmount *float32 `json:"feeAmount,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	FeeCurrency *Currency `json:"feeCurrency,omitempty"`

	// The reference for the shipment (e.g. the tracking number for a physical product, obtained from the delivery service).
	Reference *string `json:"reference,omitempty"`
}

// The URL the customer will be redirected to if the Checkout Session completed successfully. This means the Checkout succeeded, i.e. the customer authorized the order.
type SuccessUrl string

// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
type TestFlag bool

// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
type UpdatedAt int

// WebhookEndpoint defines model for webhookEndpoint.
type WebhookEndpoint struct {
	// Has the value `true` if the webhook endpoint is active and events are sent to the url specified. Has the value `false if the endpoint is inactive and the events won't be sent to the url specified.
	Active bool `json:"active"`

	// Time at which the object was created. Measured in milliseconds since the Unix epoch.
	CreatedAt CreatedAt `json:"createdAt"`

	// An optional description for your webhook endpoint.
	Description string `json:"description"`

	// The list of events to subscribe to. If not specified you will be subsribed to all events.
	EventSubscriptions *[]EventSubscription `json:"eventSubscriptions,omitempty"`

	// The unique identifier for the Webhook Endpoint object.
	Id WebhookEndpointId `json:"id"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A string representing the object’s type. The value is always `webhookEndpoint` for Webhook Endpoint objects.
	Object        string `json:"object"`
	SigningSecret string `json:"signingSecret"`

	// A flag with a value `false` if the object exists in live mode or `true` if the object exists in test mode.
	Test TestFlag `json:"test"`

	// The moment at which the object was last updated. Measured in milliseconds since the Unix epoch.
	UpdatedAt UpdatedAt `json:"updatedAt"`

	// The url which will be called when any of the events you subscribed to occur.
	Url string `json:"url"`
}

// The unique identifier for the Webhook Endpoint object.
type WebhookEndpointId string

// Expand defines model for expand.
type Expand string

// N401 defines model for 401.
type N401 struct {
	Realm      *string `json:"realm,omitempty"`
	Scheme     *string `json:"scheme,omitempty"`
	StatusCode *int    `json:"statusCode,omitempty"`
}

// WebhookEndpointNotFound defines model for WebhookEndpointNotFound.
type WebhookEndpointNotFound struct {
	Details    []interface{} `json:"details"`
	ErrorCode  string        `json:"errorCode"`
	Message    string        `json:"message"`
	StatusCode float32       `json:"statusCode"`
}

// Order
type CanceledOrder Order

// Coupons defines model for coupons.
type Coupons struct {
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

// ExpandableCheckoutSession defines model for expandableCheckoutSession.
type ExpandableCheckoutSession CheckoutSessionOptions

// ExpandableCheckoutSessions defines model for expandableCheckoutSessions.
type ExpandableCheckoutSessions struct {
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

// Expandable Order
type ExpandableOrder OrderOptions

// ExpandableOrders defines model for expandableOrders.
type ExpandableOrders struct {
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

// ExpandablePayments defines model for expandablePayments.
type ExpandablePayments struct {
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

// Expandable Refund
type ExpandableRefund RefundOptions

// ExpandableRefunds defines model for expandableRefunds.
type ExpandableRefunds struct {
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

// Payment
type MinimalPayment Payment

// Refund
type MinimalRefund Refund

// PromotionCodes defines model for promotion-codes.
type PromotionCodes struct {
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

// WebhookEndpoints defines model for webhookEndpoints.
type WebhookEndpoints struct {
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

// CheckoutSessionCreate defines model for checkout-session_create.
type CheckoutSessionCreate struct {
	// The amount intended to be collected through this order. A positive integer in the smallest currency unit.
	Amount OrderAmount `json:"amount"`

	// The URL the customer will be redirected to if the Checkout Session hasn't completed successfully. This means the Checkout failed, or the customer decided to cancel it.
	CancelUrl     CancelUrl      `json:"cancelUrl"`
	CaptureMethod *CaptureMethod `json:"captureMethod,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency Currency `json:"currency"`

	// Customer Information, the details provided here are used to pre-populate forms for your customer's checkout experiences. The more information you send, the better experience the customer will have.
	CustomerInfo CustomerInfo `json:"customerInfo"`

	// An arbitrary - ideally descriptive - long form explanation of the Order, meant to be displayed to the customer.
	Description *string `json:"description,omitempty"`

	// The line items the customer wishes to order.
	Items []Item `json:"items"`

	// Locale
	Locale *Locale `json:"locale,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A - ideally unique - string to reference the Order in your system (e.g. an order ID, etc.).
	Reference *string `json:"reference,omitempty"`

	// Shipping Information
	ShippingInfo ShippingInfo `json:"shippingInfo"`

	// The URL the customer will be redirected to if the Checkout Session completed successfully. This means the Checkout succeeded, i.e. the customer authorized the order.
	SuccessUrl SuccessUrl `json:"successUrl"`
}

// CouponCreate defines model for coupon_create.
type CouponCreate struct {
	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// required if discountType is **amount**. The amount of this coupon object.
	DiscountAmount *int32 `json:"discountAmount,omitempty"`

	// required if discountType is **percentage**. The discount percentage of this coupon object.
	DiscountPercentage *int32 `json:"discountPercentage,omitempty"`

	// Discount Type
	DiscountType DiscountType `json:"discountType"`

	// Time at which the Coupon expires. Measured in milliseconds since the Unix epoch.
	ExpiresAt *int32 `json:"expiresAt,omitempty"`

	// Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptionCount *int `json:"maxRedemptionCount,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The coupon's name, meant to be displayable to the customer.
	Name string `json:"name"`
}

// CouponUpdate defines model for coupon_update.
type CouponUpdate struct {
	// Has the value `true` if the coupon is active and can be used, or the value `false` if it is not.
	Active *bool `json:"active,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The coupon's name, meant to be displayable to the customer.
	Name *string `json:"name,omitempty"`
}

// PaymentCreate defines model for payment_create.
type PaymentCreate struct {
	// The amount to be captured.
	Amount float32 `json:"amount"`

	// Whether to cancel remaining amount in case of a partial capture. Defaults to automatic, which means only a single capture is allowed and any remaining, uncaptured amount is cancelled. Set to manual if you would like to create multiple partial captures.
	//CancelRemainder *PaymentCreateCancelRemainder `json:"cancelRemainder,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency Currency `json:"currency"`

	// An arbitrary - ideally descriptive - long form explanation of the Payment, meant to be displayed to the customer.
	Description *string `json:"description,omitempty"`

	// A list of the IDs of the Line Items of the original order this Payment captures.
	LineItems *[]LineItemId `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The unique identifier for the Payment object.
	Order OrderId `json:"order"`

	// A - ideally unique - string to reference the Payment (e.g. a customer ID, a cart ID, etc.) which can be used to reconcile the Payment with your internal systems.
	Reference *string `json:"reference,omitempty"`
}

// PromotionCodeCreate defines model for promotion-code_create.
type PromotionCodeCreate struct {
	// Has the value true (default) if the promotion code is active and can be used, or the value false if it is not.
	Active *bool `json:"active,omitempty"`

	// The customer-facing code. Regardless of case, this code must be unique across all your promotion codes.
	Code string `json:"code"`

	// The unique identifier for the Coupon object.
	Coupon CouponId `json:"coupon"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency *Currency `json:"currency,omitempty"`

	// Time at which the Promotion Code expires. Measured in milliseconds since the Unix epoch.
	ExpiresAt *int `json:"expiresAt,omitempty"`

	// A Boolean indicating if the Promotion Code should only be redeemed for customers without any successful order with the merchant. Defaults to false if not set.
	FirstTimeTransaction *bool `json:"firstTimeTransaction,omitempty"`

	// Maximum number of times this Promotion Code can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptionCount *int `json:"maxRedemptionCount,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The minimum amount required to redeem this Promotion Code (e.g., the amount of the order must be ¥10,000 or more to be applicable).
	MinimumAmount *float32 `json:"minimumAmount,omitempty"`

	// A Boolean indicating if the Promotion Code should only be redeemed once by any given customer. Defaults to false if not set.
	OnePerCustomer *bool `json:"onePerCustomer,omitempty"`
}

// PromotionCodeUpdate defines model for promotion-code_update.
type PromotionCodeUpdate struct {
	// Has the value true if the promotion code is active and can be used, or the value false if it is not.
	Active *bool `json:"active,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`
}

// RefundCreate defines model for refund_create.
type RefundCreate struct {
	// A positive amount representing how much of this payment to refund. You can refund only up to the remaining, unrefunded amount of the payment. If this is not set, the default will be the full unrefunded amount of the payment.
	Amount *float32 `json:"amount,omitempty"`

	// Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency Currency `json:"currency"`

	// An arbitrary - ideally descriptive - long form explanation of the Refund, meant to be displayed to the customer.
	Description *string `json:"description,omitempty"`

	// A list of the IDs of the Line Items of the original Payment this Refund is on.
	LineItems *[]LineItemId `json:"lineItems,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The unique identifier for the Payment object.
	Payment PaymentId `json:"payment"`

	// The reason of the Refund
	//Reason RefundCreateReason `json:"reason"`

	// A - ideally unique - string to reference the Refund which can be used to reconcile the Refund with your internal systems.
	Reference *string `json:"reference,omitempty"`
}

// RefundUpdate defines model for refund_update.
type RefundUpdate struct {
	// An arbitrary - ideally descriptive - long form explanation of the Refund, meant to be displayed to the customer.
	Description *string `json:"description,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// A - ideally unique - string to reference the Refund which can be used to reconcile the Refund with your internal systems.
	Reference *string `json:"reference,omitempty"`
}

// WebhookEndpointCreate defines model for webhook-endpoint_create.
type WebhookEndpointCreate struct {
	// An optional description for your webhook endpoint.
	Description *string `json:"description,omitempty"`

	// The list of events to subscribe to. If not specified you will be subsribed to all events.
	EventSubscriptions *[]EventSubscription `json:"eventSubscriptions,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The url which will be called when any of the events you subscribed to occur.
	Url string `json:"url"`
}

// WebhookEndpointUpdate defines model for webhook-endpoint_update.
type WebhookEndpointUpdate struct {
	// Has the value `true` if the webhook endpoint is active and events are sent to the url specified. Has the value `false if the endpoint is inactive and the events won't be sent to the url specified.
	Active *bool `json:"active,omitempty"`

	// An optional description for your webhook endpoint.
	Description *string `json:"description,omitempty"`

	// The list of events to subscribe to. If not specified you will be subsribed to all events.
	EventSubscriptions *[]EventSubscription `json:"eventSubscriptions,omitempty"`

	// Set of up to 20 key-value pairs that you can attach to the object.
	Metadata *Metadata `json:"metadata,omitempty"`

	// The url which will be called when any of the events you subscribed to occur.
	Url *string `json:"url,omitempty"`
}

// ListAllCheckoutSessionsParams defines parameters for ListAllCheckoutSessions.
type ListAllCheckoutSessionsParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`

	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *ListAllCheckoutSessionsParamsExpand `json:"expand,omitempty"`
}

// ListAllCheckoutSessionsParamsExpand defines parameters for ListAllCheckoutSessions.
type ListAllCheckoutSessionsParamsExpand string

// RetrieveACheckoutSessionParams defines parameters for RetrieveACheckoutSession.
type RetrieveACheckoutSessionParams struct {
	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *RetrieveACheckoutSessionParamsExpand `json:"expand,omitempty"`
}

// RetrieveACheckoutSessionParamsExpand defines parameters for RetrieveACheckoutSession.
type RetrieveACheckoutSessionParamsExpand string

// ListAllCouponsParams defines parameters for ListAllCoupons.
type ListAllCouponsParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`
}

// ListAllOrdersParams defines parameters for ListAllOrders.
type ListAllOrdersParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`

	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *ListAllOrdersParamsExpand `json:"expand,omitempty"`
}

// ListAllOrdersParamsExpand defines parameters for ListAllOrders.
type ListAllOrdersParamsExpand string

// RetrieveAnOrderParams defines parameters for RetrieveAnOrder.
type RetrieveAnOrderParams struct {
	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *RetrieveAnOrderParamsExpand `json:"expand,omitempty"`
}

// RetrieveAnOrderParamsExpand defines parameters for RetrieveAnOrder.
type RetrieveAnOrderParamsExpand string

// ListAllPaymentsParams defines parameters for ListAllPayments.
type ListAllPaymentsParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`

	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *ListAllPaymentsParamsExpand `json:"expand,omitempty"`
}

// ListAllPaymentsParamsExpand defines parameters for ListAllPayments.
type ListAllPaymentsParamsExpand string

// CreateAPaymentJSONBodyCancelRemainder defines parameters for CreateAPayment.
type CreateAPaymentJSONBodyCancelRemainder string

// ListAllPromotionCodesParams defines parameters for ListAllPromotionCodes.
type ListAllPromotionCodesParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`
}

// ListAllRefundsParams defines parameters for ListAllRefunds.
type ListAllRefundsParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`

	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *ListAllRefundsParamsExpand `json:"expand,omitempty"`
}

// ListAllRefundsParamsExpand defines parameters for ListAllRefunds.
type ListAllRefundsParamsExpand string

// CreateARefundJSONBodyReason defines parameters for CreateARefund.
type CreateARefundJSONBodyReason string

// RetrieveARefundParams defines parameters for RetrieveARefund.
type RetrieveARefundParams struct {
	// Set to `all` if the references within the response need to be expanded to the full objects
	Expand *RetrieveARefundParamsExpand `json:"expand,omitempty"`
}

// RetrieveARefundParamsExpand defines parameters for RetrieveARefund.
type RetrieveARefundParamsExpand string

// GetV1WebhookEndpointsParams defines parameters for GetV1WebhookEndpoints.
type GetV1WebhookEndpointsParams struct {
	// Number of objects to return. Defaults to 20 if not set.
	MaxResults *MaxResults `json:"maxResults,omitempty"`

	// The token for the page of the collection of objects.
	PageToken *PageToken `json:"pageToken,omitempty"`
}

// CreateACheckoutSessionJSONRequestBody defines body for CreateACheckoutSession for application/json ContentType.
type CreateACheckoutSessionJSONRequestBody CheckoutSessionCreate

// CreateACouponJSONRequestBody defines body for CreateACoupon for application/json ContentType.
type CreateACouponJSONRequestBody CouponCreate

// UpdateACouponJSONRequestBody defines body for UpdateACoupon for application/json ContentType.
type UpdateACouponJSONRequestBody CouponUpdate

// CreateAPaymentJSONRequestBody defines body for CreateAPayment for application/json ContentType.
type CreateAPaymentJSONRequestBody PaymentCreate

// CreateAPromotionCodeJSONRequestBody defines body for CreateAPromotionCode for application/json ContentType.
type CreateAPromotionCodeJSONRequestBody PromotionCodeCreate

// UpdateAPromotionCodeJSONRequestBody defines body for UpdateAPromotionCode for application/json ContentType.
type UpdateAPromotionCodeJSONRequestBody PromotionCodeUpdate

// CreateARefundJSONRequestBody defines body for CreateARefund for application/json ContentType.
type CreateARefundJSONRequestBody RefundCreate

// UpdateARefundJSONRequestBody defines body for UpdateARefund for application/json ContentType.
type UpdateARefundJSONRequestBody RefundUpdate

// PostV1WebhookEndpointsJSONRequestBody defines body for PostV1WebhookEndpoints for application/json ContentType.
type PostV1WebhookEndpointsJSONRequestBody WebhookEndpointCreate

// PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody defines body for PatchV1WebhookEndpointsWebhookEndpointId for application/json ContentType.
type PatchV1WebhookEndpointsWebhookEndpointIdJSONRequestBody WebhookEndpointUpdate
