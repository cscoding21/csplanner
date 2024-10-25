# Configuring Keycloak from a New Install

- Create realm __csplanner__ and import script __csplanner-realm-export.json__
- Create user __csplanner-admin__ in __master__ realm
- Create a password for csplanner-admin user by copying the appropriate secret from the secret file __kc-admin-pass.txt__ (or create a new password in the UI and change the secrets file) 
- Add all roles to __csplanner-admin__ user
- Go to __csplanner__ realm
- Go to Clients -> web-client and go to the credentials tab.
- Regenerate credentials with the button next to the __client secret__ section
- Copy new credentials into the secrets file (__kc-client-secret.txt__)
- reference: https://medium.com/@ramanamuttana/custom-attribute-in-keycloak-access-token-831b4be7384a
- run __make k8s-secrets__ to update kubernetes secret objects
