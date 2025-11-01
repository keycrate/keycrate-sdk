use serde::{Deserialize, Serialize};
use reqwest::Client;
use std::time::Duration;

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct AuthResponse {
    pub success: bool,
    pub message: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<serde_json::Value>,
}

#[derive(Debug, Clone, Default)]
pub struct AuthenticateOptions {
    pub license: Option<String>,
    pub username: Option<String>,
    pub password: Option<String>,
    pub hwid: Option<String>,
}

#[derive(Debug, Clone)]
pub struct RegisterOptions {
    pub license: String,
    pub username: String,
    pub password: String,
}

pub struct LicenseAuthClient {
    host: String,
    app_id: String,
    client: Client,
}

impl LicenseAuthClient {
    /// Create a new Keycrate client
    
    pub fn new(host: impl Into<String>, app_id: impl Into<String>) -> Self {
        let host = host.into();
        let host = if host.ends_with('/') {
            host[..host.len() - 1].to_string()
        } else {
            host
        };

        let client = Client::builder()
            .timeout(Duration::from_secs(10))
            .build()
            .unwrap_or_default();

        Self {
            host,
            app_id: app_id.into(),
            client,
        }
    }

    /// Authenticate using either a license key or username/password combination
    pub async fn authenticate(
        &self,
        opts: AuthenticateOptions,
    ) -> Result<AuthResponse, Box<dyn std::error::Error>> {
        // Validate authentication method
        if opts.license.is_none() && (opts.username.is_none() || opts.password.is_none()) {
            return Ok(AuthResponse {
                success: false,
                message: "Either license key OR (username AND password) must be provided"
                    .to_string(),
                data: None,
            });
        }

        let mut payload = serde_json::json!({
            "app_id": self.app_id,
        });

        if let Some(license) = opts.license {
            payload["license"] = serde_json::json!(license);
        }
        if let Some(username) = opts.username {
            payload["username"] = serde_json::json!(username);
        }
        if let Some(password) = opts.password {
            payload["password"] = serde_json::json!(password);
        }
        if let Some(hwid) = opts.hwid {
            payload["hwid"] = serde_json::json!(hwid);
        }

        self.make_request("/auth", payload).await
    }

    /// Register credentials for a license
    /// All parameters (license, username, password) are required
    pub async fn register(
        &self,
        opts: RegisterOptions,
    ) -> Result<AuthResponse, Box<dyn std::error::Error>> {
        // Validate all required fields
        if opts.license.is_empty() {
            return Ok(AuthResponse {
                success: false,
                message: "license is required".to_string(),
                data: None,
            });
        }
        if opts.username.is_empty() {
            return Ok(AuthResponse {
                success: false,
                message: "username is required".to_string(),
                data: None,
            });
        }
        if opts.password.is_empty() {
            return Ok(AuthResponse {
                success: false,
                message: "password is required".to_string(),
                data: None,
            });
        }

        let payload = serde_json::json!({
            "app_id": self.app_id,
            "license": opts.license,
            "username": opts.username,
            "password": opts.password,
        });

        self.make_request("/register", payload).await
    }

    async fn make_request(
        &self,
        endpoint: &str,
        payload: serde_json::Value,
    ) -> Result<AuthResponse, Box<dyn std::error::Error>> {
        let url = format!("{}{}", self.host, endpoint);

        let response = self
            .client
            .post(&url)
            .header("Content-Type", "application/json")
            .json(&payload)
            .send()
            .await?;

        let result: AuthResponse = response.json().await?;
        Ok(result)
    }
}

