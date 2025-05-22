# Configuring Keycloak from a New Install

Local URL: __http://localhost:8080/auth__

- Create user __csplanner-admin__ in __master__ realm
    - Create a password for csplanner-admin user by copying the appropriate secret from the secret file __kc-admin-pass.txt__ (or create a new password in the UI and change the secrets file) 
    - Add all roles to __csplanner-admin__ user
- Create client __cs-provisioner__ in __master__ realm
    - check client authentication / Standard flow - Direct access grants - Service accounts roles
    - assign all service account roles
    - copy client-secret from credentials to CSP_SECURITY_KEYCLOAKPROVISIONERCLIENTSECRET in .env.local
- reference: https://medium.com/@ramanamuttana/custom-attribute-in-keycloak-access-token-831b4be7384a
- run __make k8s-secrets__ to update kubernetes secret objects
