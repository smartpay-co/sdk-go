package smartpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/eknkc/basex"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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

// CalculateWebhookSignatureMiddleware return a middleware that adds Calculated-Signature to request header
func CalculateWebhookSignatureMiddleware(signingSecret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signature := r.Header.Get("Smartpay-Signature")
		signatureTimestamp := r.Header.Get("Smartpay-Signature-Timestamp")

		if signature != "" && signatureTimestamp != "" {
			// Read body and reassign a new body reader for the next middleware
			bodyBytes, _ := io.ReadAll(r.Body)
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			data := signatureTimestamp + "." + string(bodyBytes)
			sha, err := CalculateSignature(signingSecret, data)
			if err != nil {
				log.Println("calculateSignature failed: ", err)
			}
			r.Header.Set("Calculated-Signature", sha)
		}

		next.ServeHTTP(w, r)
	})
}

func CalculateSignature(signingSecret string, data string) (sha string, err error) {
	enc, err := basex.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	if err != nil {
		return "", err
	}

	secret, err := enc.Decode(signingSecret)
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	sha = hex.EncodeToString(h.Sum(nil))

	return
}

func VerifyWebhookSignature(signingSecret string, data string, signature string) bool {
	sha, _ := CalculateSignature(signingSecret, data)
	return sha == signature
}

func newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Add Idempotency-Key for POST, PUT and PATCH requests.
	// ref: https://en.docs.smartpay.co/reference/idempotency
	method = strings.ToUpper(method)
	if method == "POST" || method == "PUT" || method == "PATCH" {
		req.Header.Set("Idempotency-Key", uuid.New().String())
	}

	return req, nil
}
