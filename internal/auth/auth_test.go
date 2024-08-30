package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// given
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey my-api-key")

	// when
	got, err := GetAPIKey(headers)

	// then
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := "my-api-key"
	if got != want {
		t.Errorf("expected API key to be %s, but got %s", want, got)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	// given
	headers := make(http.Header)

	// when
	_, err := GetAPIKey(headers)

	// then
	if err == nil {
		t.Error("expected error, but got nil")
	}

	expectedError := ErrNoAuthHeaderIncluded
	if err != expectedError {
		t.Errorf("expected error to be %v, but got %v", expectedError, err)
	}
}
