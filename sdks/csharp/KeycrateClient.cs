using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace Keycrate
{
    public class AuthenticateOptions
    {
        public string License { get; set; }
        public string Username { get; set; }
        public string Password { get; set; }
        public string Hwid { get; set; }
    }

    public class RegisterOptions
    {
        public string License { get; set; }
        public string Username { get; set; }
        public string Password { get; set; }
    }

    public class KeycrateClient
    {
        private readonly string _host;
        private readonly string _appId;
        private readonly HttpClient _httpClient;
        private readonly JsonSerializerOptions _jsonOptions;

        public KeycrateClient(string host, string appId)
        {
            _host = host.TrimEnd('/');
            _appId = appId;
            _httpClient = new HttpClient { Timeout = TimeSpan.FromSeconds(15) };
            _jsonOptions = new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            };
        }

        public async Task<Dictionary<string, object>> AuthenticateAsync(AuthenticateOptions options)
        {
            if (string.IsNullOrEmpty(options.License) &&
                (string.IsNullOrEmpty(options.Username) || string.IsNullOrEmpty(options.Password)))
            {
                return new Dictionary<string, object>
                {
                    ["success"] = false,
                    ["message"] = "Either license or (username + password) required"
                };
            }

            var payload = new Dictionary<string, object> { ["app_id"] = _appId };
            if (!string.IsNullOrEmpty(options.License)) payload["license"] = options.License;
            if (!string.IsNullOrEmpty(options.Username)) payload["username"] = options.Username;
            if (!string.IsNullOrEmpty(options.Password)) payload["password"] = options.Password;
            if (!string.IsNullOrEmpty(options.Hwid)) payload["hwid"] = options.Hwid;

            return await MakeRequestAsync("/auth", payload);
        }

        public async Task<Dictionary<string, object>> RegisterAsync(RegisterOptions options)
        {
            if (string.IsNullOrEmpty(options.License))
                return new() { ["success"] = false, ["message"] = "license required" };
            if (string.IsNullOrEmpty(options.Username))
                return new() { ["success"] = false, ["message"] = "username required" };
            if (string.IsNullOrEmpty(options.Password))
                return new() { ["success"] = false, ["message"] = "password required" };

            var payload = new Dictionary<string, object>
            {
                ["app_id"] = _appId,
                ["license"] = options.License,
                ["username"] = options.Username,
                ["password"] = options.Password
            };

            return await MakeRequestAsync("/register", payload);
        }

        private async Task<Dictionary<string, object>> MakeRequestAsync(string endpoint, Dictionary<string, object> payload)
        {
            try
            {
                var content = new StringContent(
                    JsonSerializer.Serialize(payload, _jsonOptions),
                    Encoding.UTF8,
                    "application/json");

                var response = await _httpClient.PostAsync(_host + endpoint, content);
                string body = await response.Content.ReadAsStringAsync();

                var result = JsonSerializer.Deserialize<Dictionary<string, object>>(body, _jsonOptions);

                return result ?? new()
                {
                    ["success"] = false,
                    ["message"] = "Failed to parse response"
                };
            }
            catch (Exception ex)
            {
                return new()
                {
                    ["success"] = false,
                    ["message"] = $"Request failed: {ex.Message}"
                };
            }
        }
    }
}
