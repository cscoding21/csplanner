package provision

import (
	"context"
	"csserver/internal/config"
	"csserver/internal/utils"
	"encoding/json"

	"github.com/Nerzal/gocloak/v13"
	"github.com/cscoding21/csgen"

	log "github.com/sirupsen/logrus"
)

type RealmVals struct {
	RealmID      string
	WebClientID  string
	OrgKey       string
	ClientSecret string
}

func getKeycloakClient() *gocloak.GoCloak {
	return gocloak.NewClient(config.Config.Security.KeycloakURL)
}

// CreateNewOrgRealm create a new realm to contain org users
func CreateNewOrgRealm(ctx context.Context, name string) (string, error) {
	newRealm, err := getNewRealmRepresentation(name, GenerateOrgKey(name), config.Config.Security.KeycloakClientSecret)
	if err != nil {
		log.Fatal(err)
	}

	client := getKeycloakClient()
	token, err := client.LoginClient(ctx,
		config.Config.Security.KeycloakProvisionerClientID,
		config.Config.Security.KeycloakProvisionerClientSecret,
		config.Config.Security.KeycloakMasterRealmName)
	if err != nil {
		log.Fatal(err)
	}

	return client.CreateRealm(ctx, token.AccessToken, newRealm)
}

// DeleteOrgRealm delete a realm from Keycloak
func DeleteOrgRealm(ctx context.Context, name string) error {
	realmName := GenerateOrgKey(name)

	client := getKeycloakClient()
	token, err := client.LoginClient(ctx,
		config.Config.Security.KeycloakProvisionerClientID,
		config.Config.Security.KeycloakProvisionerClientSecret,
		config.Config.Security.KeycloakMasterRealmName)
	if err != nil {
		log.Fatal(err)
	}

	err = client.DeleteRealm(ctx, token.AccessToken, realmName)
	if err != nil {
		log.Errorf("ERROR: %s\n", err)
	}

	return err
}

func getNewRealmRepresentation(name string, orgKey string, clientSecret string) (gocloak.RealmRepresentation, error) {
	var realmRep gocloak.RealmRepresentation

	rv := RealmVals{
		RealmID:      utils.GenerateBase64UUID(),
		WebClientID:  utils.GenerateBase64UUID(),
		OrgKey:       orgKey,
		ClientSecret: clientSecret,
	}

	realmJSON := csgen.ExecuteTemplate("newRealm", realmTemplate, rv)

	err := json.Unmarshal([]byte(realmJSON), &realmRep)
	if err != nil {
		return realmRep, err
	}

	realmRep.DisplayName = &name

	return realmRep, nil
}

var realmTemplate = `
{
    "id": "{{.RealmID}}",
    "realm": "{{.OrgKey}}",
    "enabled": true,
    "notBefore": 0,
    "defaultSignatureAlgorithm": "RS256",
    "clients": [
      {
        "clientId": "web-client",
        "name": "Web API Login",
        "enabled": true,
        "alwaysDisplayInConsole": true,
        "clientAuthenticatorType": "client-secret",
        "secret": "{{.ClientSecret}}",
        "redirectUris": [
          "/*"
        ],
        "webOrigins": [
          "/*"
        ],
        "notBefore": 0,
        "bearerOnly": false,
        "consentRequired": false,
        "standardFlowEnabled": true,
        "implicitFlowEnabled": false,
        "directAccessGrantsEnabled": true,
        "serviceAccountsEnabled": true,
        "publicClient": false,
        "frontchannelLogout": true,
        "protocol": "openid-connect",
        "attributes": {
          "oidc.ciba.grant.enabled": "false",
          "client.secret.creation.time": "1726195662",
          "backchannel.logout.session.required": "true",
          "display.on.consent.screen": "false",
          "oauth2.device.authorization.grant.enabled": "false",
          "backchannel.logout.revoke.offline.tokens": "false"
        },
        "authenticationFlowBindingOverrides": {},
        "fullScopeAllowed": true,
        "nodeReRegistrationTimeout": -1,
        "protocolMappers": [
          {
            "name": "dbid",
            "protocol": "openid-connect",
            "protocolMapper": "oidc-usermodel-attribute-mapper",
            "consentRequired": false,
            "config": {
              "introspection.token.claim": "true",
              "userinfo.token.claim": "true",
              "user.attribute": "dbid",
              "id.token.claim": "true",
              "lightweight.claim": "false",
              "access.token.claim": "true",
              "claim.name": "dbid",
              "jsonType.label": "String"
            }
          },
          {
            "name": "profileImage",
            "protocol": "openid-connect",
            "protocolMapper": "oidc-usermodel-attribute-mapper",
            "consentRequired": false,
            "config": {
              "introspection.token.claim": "true",
              "userinfo.token.claim": "true",
              "user.attribute": "profileImage",
              "id.token.claim": "true",
              "lightweight.claim": "false",
              "access.token.claim": "true",
              "claim.name": "profileImage",
              "jsonType.label": "String"
            }
          }
        ]
      }
    ]   
  }
`
