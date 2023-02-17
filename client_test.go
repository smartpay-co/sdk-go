package smartpay

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRetryableClient(t *testing.T) {
	var requestCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	retryingClient, _ := NewClientWithResponses("secretKey", "publicKey",
		WithBaseURL(server.URL),
		WithRetryOptions(2, time.Second, 5*time.Second),
	)

	ctx := context.Background()
	resp, err := retryingClient.RetrieveAToken(ctx, "tokenId")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if requestCount != 3 {
		t.Errorf("expected 3 requests, got %d", requestCount)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestHttpClientWithIdempotencyHeader(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the "Idempotency-Key" header is present in the request
		idempotencyKey := r.Header.Get("Idempotency-Key")
		if idempotencyKey == "" {
			t.Error("Idempotency-Key header not found in the request")
		}
	}))
	defer ts.Close()

	client, _ := NewClientWithResponses("secretKey", "publicKey", WithBaseURL(ts.URL))

	ctx := context.Background()
	payload := CreateAWebhookEndpointJSONRequestBody{
		EventSubscriptions: Ptr([]EventSubscription{
			EventSubscriptionMerchantUserCreated,
			EventSubscriptionPromotionCodeUpdated,
		}),
		Url: "https://www.example.com",
	}
	resp, err := client.CreateAWebhookEndpoint(ctx, payload)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestRetryableHttpClientWithIdempotencyHeader(t *testing.T) {
	// Create a test server
	var requestCount int
	var lastIdempotencyKey string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the "Idempotency-Key" header is present in the request
		requestCount++
		idempotencyKey := r.Header.Get("Idempotency-Key")

		if idempotencyKey == "" {
			t.Error("Idempotency-Key header not found in the request")
		}

		// Assign the current idempotency key for comparison
		if lastIdempotencyKey == "" {
			lastIdempotencyKey = idempotencyKey
		} else {
			if lastIdempotencyKey != idempotencyKey {
				t.Error("Idempotency-Key header changed in the same request")
			}
		}
		if requestCount == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer ts.Close()

	client, _ := NewClientWithResponses("secretKey", "publicKey", WithBaseURL(ts.URL))

	ctx := context.Background()
	payload := CreateAWebhookEndpointJSONRequestBody{
		EventSubscriptions: Ptr([]EventSubscription{
			EventSubscriptionMerchantUserCreated,
			EventSubscriptionPromotionCodeUpdated,
		}),
		Url: "https://www.example.com",
	}
	resp, err := client.CreateAWebhookEndpoint(ctx, payload)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestHttpClientWithCustomIdempotencyHeader(t *testing.T) {
	idempotencyKey := "12345678"
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the "Idempotency-Key" header is present in the request
		receivedIdempotencyKey := r.Header.Get("Idempotency-Key")
		if idempotencyKey != receivedIdempotencyKey {
			t.Error("Idempotency-Key header not properly set")
		}
	}))
	defer ts.Close()

	client, _ := NewClientWithResponses("secretKey", "publicKey", WithBaseURL(ts.URL))

	ctx := context.Background()
	payload := CreateAWebhookEndpointJSONRequestBody{
		EventSubscriptions: Ptr([]EventSubscription{
			EventSubscriptionMerchantUserCreated,
			EventSubscriptionPromotionCodeUpdated,
		}),
		Url: "https://www.example.com",
	}
	resp, _ := client.CreateAWebhookEndpoint(ctx, payload, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Idempotency-Key", idempotencyKey)
		return nil
	})
	if resp != nil {
		defer resp.Body.Close()
	}
}
