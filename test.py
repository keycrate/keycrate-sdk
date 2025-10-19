import keycrate


def main():
    host = input("Enter host (e.g., http://127.0.0.1:8787): ").strip()
    if host == "":
        host = "http://127.0.0.1:8787"

    app_id = input("Enter app_id: ").strip()
    if app_id == "":
        app_id = "57d87dfa-18a6-4eed-9074-f37418067c47"

    client = keycrate.configurate(host=host, app_id=app_id)

    action = input("Do you want to (1) Authenticate or (2) Register? [1/2]: ").strip()

    if action == "1":
        # AUTHENTICATE
        use_hwid = input("Do you want to use HWID? [y/n]: ").strip().lower() == "y"
        hwid = input("Enter HWID: ").strip() if use_hwid else None

        auth_type = input("Authenticate with (1) License or (2) Username/Password? [1/2]: ").strip()
        
        if auth_type == "1":
            # License-based authentication
            license_key = input("Enter license key: ").strip()
            resp = client.authenticate(license=license_key, hwid=hwid)
        else:
            # Username/password authentication
            username = input("Enter username: ").strip()
            password = input("Enter password: ").strip()
            resp = client.authenticate(username=username, password=password, hwid=hwid)

        print("\nAuthentication response:")
        print(resp)

    elif action == "2":
        # REGISTER
        license_key = input("Enter license key to register: ").strip()
        username = input("Enter desired username: ").strip()
        password = input("Enter password: ").strip()
        
        resp = client.register(license=license_key, username=username, password=password)
        
        print("\nRegistration response:")
        print(resp)

    else:
        print("Invalid option. Exiting.")


if __name__ == "__main__":
    main()