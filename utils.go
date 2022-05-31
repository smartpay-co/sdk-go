package smartpay

import (
	"encoding/json"
	"net/url"
)

// Ptr Inline pointer helper
func Ptr[T any](v T) *T {
	return &v
}

func ConvertToStruct[T any](from interface{}) (to T, err error) {
	jsonData, err := json.Marshal(from)
	if err != nil {
		return to, err
	}
	err = json.Unmarshal(jsonData, &to)
	if err != nil {
		return to, err
	}
	return
}

func (checkoutSessionUrl *CheckoutSessioUrl) WithPromotionCode(promotionCode string) (checkoutSessionUrlWithPromotionCode string, err error) {
	rawUrl, err := url.Parse(string(*checkoutSessionUrl))
	if err != nil {
		return
	}
	values := rawUrl.Query()
	values.Add("promotion-code", promotionCode)
	rawUrl.RawQuery = values.Encode()
	checkoutSessionUrlWithPromotionCode = rawUrl.String()
	return
}
