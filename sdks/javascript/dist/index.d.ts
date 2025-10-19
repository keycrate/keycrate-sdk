/**
 * Keycrate License Authentication SDK
 * JavaScript/TypeScript implementation
 */
interface AuthResponse {
    success: boolean;
    message: string;
    data?: Record<string, unknown>;
}
interface AuthenticateOptions {
    license?: string;
    username?: string;
    password?: string;
    hwid?: string;
}
interface RegisterOptions {
    license: string;
    username: string;
    password: string;
}
declare class LicenseAuthClient {
    private host;
    private appId;
    private timeout;
    /**
     * Initialize the client
     * @param host - Base URL of the API (e.g., 'https://api.example.com')
     * @param appId - Application ID for authentication
     * @param timeout - Request timeout in milliseconds (default: 10000)
     */
    constructor(host: string, appId: string, timeout?: number);
    /**
     * Authenticate using either a license key or username/password combination
     * @param options - Authentication options
     * @returns Authentication response with success status and message
     */
    authenticate(options: AuthenticateOptions): Promise<AuthResponse>;
    /**
     * Register credentials for a license
     * @param options - Registration options (all required)
     * @returns Registration response with success status and message
     */
    register(options: RegisterOptions): Promise<AuthResponse>;
    /**
     * Internal method to make HTTP requests
     */
    private _makeRequest;
}
/**
 * Factory function to create and configure a client
 * @param host - Base URL of the API
 * @param appId - Application ID
 * @returns Configured LicenseAuthClient instance
 */
declare function configurate(host: string, appId: string): LicenseAuthClient;
export { LicenseAuthClient, configurate, AuthResponse, AuthenticateOptions, RegisterOptions };
//# sourceMappingURL=index.d.ts.map