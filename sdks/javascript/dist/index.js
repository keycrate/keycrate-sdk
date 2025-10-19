/**
 * Keycrate License Authentication SDK
 * JavaScript/TypeScript implementation
 */
class LicenseAuthClient {
    /**
     * Initialize the client
     * @param host - Base URL of the API (e.g., 'https://api.example.com')
     * @param appId - Application ID for authentication
     * @param timeout - Request timeout in milliseconds (default: 10000)
     */
    constructor(host, appId, timeout = 10000) {
        this.host = host.replace(/\/$/, ''); // Remove trailing slash
        this.appId = appId;
        this.timeout = timeout;
    }
    /**
     * Authenticate using either a license key or username/password combination
     * @param options - Authentication options
     * @returns Authentication response with success status and message
     */
    async authenticate(options) {
        const { license, username, password, hwid } = options;
        // Validate authentication method
        if (!license && !(username && password)) {
            return {
                success: false,
                message: 'Either license key OR (username AND password) must be provided'
            };
        }
        const payload = {
            app_id: this.appId
        };
        if (license)
            payload.license = license;
        if (username)
            payload.username = username;
        if (password)
            payload.password = password;
        if (hwid)
            payload.hwid = hwid;
        try {
            const response = await this._makeRequest('/auth', payload);
            return response;
        }
        catch (error) {
            return {
                success: false,
                message: `Request failed: ${error instanceof Error ? error.message : String(error)}`
            };
        }
    }
    /**
     * Register credentials for a license
     * @param options - Registration options (all required)
     * @returns Registration response with success status and message
     */
    async register(options) {
        const { license, username, password } = options;
        // Validate all required fields
        if (!license) {
            return { success: false, message: 'license is required' };
        }
        if (!username) {
            return { success: false, message: 'username is required' };
        }
        if (!password) {
            return { success: false, message: 'password is required' };
        }
        const payload = {
            app_id: this.appId,
            license,
            username,
            password
        };
        try {
            const response = await this._makeRequest('/register', payload);
            return response;
        }
        catch (error) {
            return {
                success: false,
                message: `Request failed: ${error instanceof Error ? error.message : String(error)}`
            };
        }
    }
    /**
     * Internal method to make HTTP requests
     */
    async _makeRequest(endpoint, payload) {
        const url = `${this.host}${endpoint}`;
        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), this.timeout);
        try {
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload),
                signal: controller.signal
            });
            const data = await response.json();
            return data;
        }
        catch (error) {
            if (error instanceof Error && error.name === 'AbortError') {
                throw new Error(`Request timeout after ${this.timeout}ms`);
            }
            throw error;
        }
        finally {
            clearTimeout(timeoutId);
        }
    }
}
/**
 * Factory function to create and configure a client
 * @param host - Base URL of the API
 * @param appId - Application ID
 * @returns Configured LicenseAuthClient instance
 */
function configurate(host, appId) {
    return new LicenseAuthClient(host, appId);
}
export { LicenseAuthClient, configurate };
