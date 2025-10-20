package keycrate

import (
	"testing"
)

func TestAuthenticateValidation(t *testing.T) {
	client := New("http://127.0.0.1:8787", "test-app-id")

	// Test: Missing both authentication methods
	result, err := client.Authenticate(AuthenticateOptions{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Success {
		t.Error("Expected failure for missing auth method")
	}
	if result.Message != "Either license key OR (username AND password) must be provided" {
		t.Errorf("Expected validation message, got: %s", result.Message)
	}

	// Test: Only username, no password
	result, err = client.Authenticate(AuthenticateOptions{
		Username: "testuser",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Success {
		t.Error("Expected failure for missing password")
	}

	// Test: Only password, no username
	result, err = client.Authenticate(AuthenticateOptions{
		Password: "testpass",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Success {
		t.Error("Expected failure for missing username")
	}
}

func TestRegisterValidation(t *testing.T) {
	client := New("http://127.0.0.1:8787", "test-app-id")

	tests := []struct {
		name    string
		opts    RegisterOptions
		wantErr string
	}{
		{
			name:    "Missing license",
			opts:    RegisterOptions{Username: "user", Password: "pass"},
			wantErr: "license is required",
		},
		{
			name:    "Missing username",
			opts:    RegisterOptions{License: "lic", Password: "pass"},
			wantErr: "username is required",
		},
		{
			name:    "Missing password",
			opts:    RegisterOptions{License: "lic", Username: "user"},
			wantErr: "password is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.Register(tt.opts)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if result.Success {
				t.Error("Expected failure for missing field")
			}
			if result.Message != tt.wantErr {
				t.Errorf("Expected %q, got %q", tt.wantErr, result.Message)
			}
		})
	}
}

func TestAuthenticateLicense(t *testing.T) {
	// This test requires a running API server
	// Uncomment to test against live server
	/*
	client := New("http://127.0.0.1:8787", "57d87dfa-18a6-4eed-9074-f37418067c47")

	result, err := client.Authenticate(AuthenticateOptions{
		License: "test-license-key",
	})
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	t.Logf("Response: %+v", result)
	*/
}

func TestRegister(t *testing.T) {
	// This test requires a running API server
	// Uncomment to test against live server
	/*
	client := New("http://127.0.0.1:8787", "57d87dfa-18a6-4eed-9074-f37418067c47")

	result, err := client.Register(RegisterOptions{
		License:  "test-license",
		Username: "testuser",
		Password: "testpass",
	})
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	t.Logf("Response: %+v", result)
	*/
}