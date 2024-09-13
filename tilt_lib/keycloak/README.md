# Configuring Keycloak from a New Install

- Create realm __csplanner__
- Create user __csplanner-admin__ in __master__ realm
- Create a password for csplanner-admin user by copying the appropriate secret from the secret file 
- Add all roles to __csplanner-admin__ user
- Go to __csplanner__ realm
- Add __web-client__ client to realm (make sure to check confidential)
- Copy password in credentials tab and paste into secret file
- In realm settings, go to the __users__ tab
- Add custom attribute __profileImage__
- Go to client -> client-scopes
- Add a mapper to token for __profileImage__ attribute
- reference: https://medium.com/@ramanamuttana/custom-attribute-in-keycloak-access-token-831b4be7384a
- run __make k8s-secrets__ to update kubernetes secret objects
