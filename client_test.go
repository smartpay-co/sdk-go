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
