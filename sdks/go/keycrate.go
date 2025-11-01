package keycrate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// AuthResponse represents the API response structure
type AuthResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// AuthenticateOptions holds authentication parameters
type AuthenticateOptions struct {
	License  string
	Username string
	Password string
	HWID     string
}

// RegisterOptions holds registration parameters
type RegisterOptions struct {
	License  string
	Username string
	Password string
}

// Client represents a Keycrate license client
type Client struct {
	host    string
	appID   string
	timeout time.Duration
	client  *http.Client
}

// New creates a new Keycrate client
// host: Base URL of the API (e.g., "https://api.example.com")
// appID: Application ID for authentication
func New(host, appID string) *Client {
	host = strings.TrimSuffix(host, "/")
	
	return &Client{
		host:    host,
		appID:   appID,
		timeout: 10 * time.Second,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Authenticate authenticates using either a license key or username/password combination
// Returns AuthResponse with success status and message
func (c *Client) Authenticate(opts AuthenticateOptions) (*AuthResponse, error) {
	// Validate authentication method
	if opts.License == "" && (opts.Username == "" || opts.Password == "") {
		return &AuthResponse{
			Success: false,
			Message: "Either license key OR (username AND password) must be provided",
		}, nil
	}

	payload := map[string]string{
		"app_id": c.appID,
	}

	if opts.License != "" {
		payload["license"] = opts.License
	}
	if opts.Username != "" {
		payload["username"] = opts.Username
	}
	if opts.Password != "" {
		payload["password"] = opts.Password
	}
	if opts.HWID != "" {
		payload["hwid"] = opts.HWID
	}

	return c.makeRequest("/auth", payload)
}

// Register registers credentials for a license
// All parameters (license, username, password) are required
// Returns AuthResponse with success status and message
func (c *Client) Register(opts RegisterOptions) (*AuthResponse, error) {
	// Validate all required fields
	if opts.License == "" {
		return &AuthResponse{
			Success: false,
			Message: "license is required",
		}, nil
	}
	if opts.Username == "" {
		return &AuthResponse{
			Success: false,
			Message: "username is required",
		}, nil
	}
	if opts.Password == "" {
		return &AuthResponse{
			Success: false,
			Message: "password is required",
		}, nil
	}

	payload := map[string]string{
		"app_id":   c.appID,
		"license":  opts.License,
		"username": opts.Username,
		"password": opts.Password,
	}

	return c.makeRequest("/register", payload)
}

// makeRequest makes an HTTP POST request to the API
func (c *Client) makeRequest(endpoint string, payload map[string]string) (*AuthResponse, error) {
	url := c.host + endpoint

	// Marshal payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Make request
	resp, err := c.client.Do(req)
	if err != nil {
		return &AuthResponse{
			Success: false,
			Message: fmt.Sprintf("Request failed: %v", err),
		}, nil
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Unmarshal response
	var result AuthResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return &AuthResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid response from server (HTTP %d)", resp.StatusCode),
		}, nil
	}

	return &result, nil
}