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
	RealmID     string
	WebClientID string
	OrgKey      string
}

func getKeycloakClient() *gocloak.GoCloak {
	return gocloak.NewClient(config.Config.Security.KeycloakURL)
}

// CreateNewOrgRealm create a new realm to contain org users
func CreateNewOrgRealm(ctx context.Context, name string) (string, error) {
	newRealm, err := getNewRealmRepresentation(name, GenerateOrgKey(name))
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

func getNewRealmRepresentation(name string, orgKey string) (gocloak.RealmRepresentation, error) {
	var realmRep gocloak.RealmRepresentation

	rv := RealmVals{}
	rv.RealmID = utils.GenerateBase64UUID()
	rv.WebClientID = utils.GenerateBase64UUID()
	rv.OrgKey = orgKey

	realmJSON := csgen.ExecuteTemplate("newRealm", realmTemplate, rv)

	err := json.Unmarshal([]byte(realmJSON), &realmRep)
	if err != nil {
		return realmRep, err
	}

	realmRep.DisplayName = &name

	return realmRep, nil
}

/*
var realmRolesForUser = []gocloak.Role{
	{Name: utils.ValToRef("view-clients")},
	{Name: utils.ValToRef("query-groups")},
	{Name: utils.ValToRef("view-realm")},
	{Name: utils.ValToRef("impersonation")},
	{Name: utils.ValToRef("manage-realm")},
	{Name: utils.ValToRef("create-client")},
	{Name: utils.ValToRef("manage-events")},
	{Name: utils.ValToRef("view-identity-providers")},
	{Name: utils.ValToRef("manage-authorization")},
	{Name: utils.ValToRef("view-users")},
	{Name: utils.ValToRef("query-realms")},
	{Name: utils.ValToRef("query-users")},
	{Name: utils.ValToRef("manage-clients")},
	{Name: utils.ValToRef("query-clients")},
	{Name: utils.ValToRef("manage-users")},
	{Name: utils.ValToRef("manage-identity-providers")},
	{Name: utils.ValToRef("view-events")},
	{Name: utils.ValToRef("view-authorization")},
}
*/

var realmTemplate = `{
    "id": "{{.RealmID}}",
    "realm": "{{.OrgKey}}",
    "notBefore": 0,
    "defaultSignatureAlgorithm": "RS256",
    "revokeRefreshToken": true,
    "refreshTokenMaxReuse": 0,
    "accessTokenLifespan": 600,
    "accessTokenLifespanForImplicitFlow": 900,
    "ssoSessionIdleTimeout": 1800,
    "ssoSessionMaxLifespan": 36000,
    "ssoSessionIdleTimeoutRememberMe": 0,
    "ssoSessionMaxLifespanRememberMe": 0,
    "offlineSessionIdleTimeout": 2592000,
    "offlineSessionMaxLifespanEnabled": false,
    "offlineSessionMaxLifespan": 5184000,
    "clientSessionIdleTimeout": 0,
    "clientSessionMaxLifespan": 0,
    "clientOfflineSessionIdleTimeout": 0,
    "clientOfflineSessionMaxLifespan": 0,
    "accessCodeLifespan": 60,
    "accessCodeLifespanUserAction": 300,
    "accessCodeLifespanLogin": 1800,
    "actionTokenGeneratedByAdminLifespan": 43200,
    "actionTokenGeneratedByUserLifespan": 300,
    "oauth2DeviceCodeLifespan": 600,
    "oauth2DevicePollingInterval": 5,
    "enabled": true,
    "sslRequired": "external",
    "registrationAllowed": false,
    "registrationEmailAsUsername": false,
    "rememberMe": false,
    "verifyEmail": false,
    "loginWithEmailAllowed": true,
    "duplicateEmailsAllowed": false,
    "resetPasswordAllowed": false,
    "editUsernameAllowed": false,
    "bruteForceProtected": false,
    "permanentLockout": false,
    "maxTemporaryLockouts": 0,
    "maxFailureWaitSeconds": 900,
    "minimumQuickLoginWaitSeconds": 60,
    "waitIncrementSeconds": 60,
    "quickLoginCheckMilliSeconds": 1000,
    "maxDeltaTimeSeconds": 43200,
    "failureFactor": 30,
    "roles": {
        "client": {
            "realm-management": [],
            "web-client": [],
            "security-admin-console": [],
            "admin-cli": [],
            "account-console": [],
            "broker": [],
            "account": []
        }
    },
    "groups": [],
    "requiredCredentials": [
        "password"
    ],
    "otpPolicyType": "totp",
    "otpPolicyAlgorithm": "HmacSHA1",
    "otpPolicyInitialCounter": 0,
    "otpPolicyDigits": 6,
    "otpPolicyLookAheadWindow": 1,
    "otpPolicyPeriod": 30,
    "otpPolicyCodeReusable": false,
    "otpSupportedApplications": [
        "totpAppFreeOTPName",
        "totpAppGoogleName",
        "totpAppMicrosoftAuthenticatorName"
    ],
    "localizationTexts": {},
    "webAuthnPolicyRpEntityName": "keycloak",
    "webAuthnPolicySignatureAlgorithms": [
        "ES256"
    ],
    "webAuthnPolicyRpId": "",
    "webAuthnPolicyAttestationConveyancePreference": "not specified",
    "webAuthnPolicyAuthenticatorAttachment": "not specified",
    "webAuthnPolicyRequireResidentKey": "not specified",
    "webAuthnPolicyUserVerificationRequirement": "not specified",
    "webAuthnPolicyCreateTimeout": 0,
    "webAuthnPolicyAvoidSameAuthenticatorRegister": false,
    "webAuthnPolicyAcceptableAaguids": [],
    "webAuthnPolicyExtraOrigins": [],
    "webAuthnPolicyPasswordlessRpEntityName": "keycloak",
    "webAuthnPolicyPasswordlessSignatureAlgorithms": [
        "ES256"
    ],
    "webAuthnPolicyPasswordlessRpId": "",
    "webAuthnPolicyPasswordlessAttestationConveyancePreference": "not specified",
    "webAuthnPolicyPasswordlessAuthenticatorAttachment": "not specified",
    "webAuthnPolicyPasswordlessRequireResidentKey": "not specified",
    "webAuthnPolicyPasswordlessUserVerificationRequirement": "not specified",
    "webAuthnPolicyPasswordlessCreateTimeout": 0,
    "webAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister": false,
    "webAuthnPolicyPasswordlessAcceptableAaguids": [],
    "webAuthnPolicyPasswordlessExtraOrigins": [],
    "users": [
        {
            "username": "service-account-web-client",
            "emailVerified": false,
            "createdTimestamp": 1726195662219,
            "enabled": true,
            "totp": false,
            "serviceAccountClientId": "web-client",
            "disableableCredentialTypes": [],
            "requiredActions": [],
            "realmRoles": [
                "default-roles-csplanner"
            ],
            "notBefore": 0,
            "groups": []
        }
    ],
    "clients": [
        {
            "id": "{{.WebClientID}}",
            "clientId": "web-client",
            "name": "Web API Login",
            "description": "",
            "rootUrl": "",
            "adminUrl": "",
            "baseUrl": "",
            "surrogateAuthRequired": false,
            "enabled": true,
            "alwaysDisplayInConsole": false,
            "clientAuthenticatorType": "client-secret",
            "secret": "**********",
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
                    "name": "Client ID",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usersessionmodel-note-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.session.note": "client_id",
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "client_id",
                        "jsonType.label": "String"
                    }
                },
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
                    "name": "Client Host",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usersessionmodel-note-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.session.note": "clientHost",
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "clientHost",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "Client IP Address",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usersessionmodel-note-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.session.note": "clientAddress",
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "clientAddress",
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
            ],
            "defaultClientScopes": [
                "web-origins",
                "acr",
                "roles",
                "profile",
                "basic",
                "email"
            ],
            "optionalClientScopes": [
                "address",
                "phone",
                "offline_access",
                "microprofile-jwt"
            ]
        }
    ],
    "clientScopes": [
        {
            "name": "acr",
            "description": "OpenID Connect scope for add acr (authentication context class reference) to the token",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "false",
                "display.on.consent.screen": "false"
            },
            "protocolMappers": [
                {
                    "name": "acr loa level",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-acr-mapper",
                    "consentRequired": false,
                    "config": {
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true"
                    }
                }
            ]
        },
        {
            "name": "web-origins",
            "description": "OpenID Connect scope for add allowed web origins to the access token",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "false",
                "consent.screen.text": "",
                "display.on.consent.screen": "false"
            },
            "protocolMappers": [
                {
                    "name": "allowed web origins",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-allowed-origins-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "access.token.claim": "true"
                    }
                }
            ]
        },
        {
            "name": "offline_access",
            "description": "OpenID Connect built-in scope: offline_access",
            "protocol": "openid-connect",
            "attributes": {
                "consent.screen.text": "${offlineAccessScopeConsentText}",
                "display.on.consent.screen": "true"
            }
        },
        {
            "name": "phone",
            "description": "OpenID Connect built-in scope: phone",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "true",
                "consent.screen.text": "${phoneScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "phone number",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "phoneNumber",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "phone_number",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "phone number verified",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "phoneNumberVerified",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "phone_number_verified",
                        "jsonType.label": "boolean"
                    }
                }
            ]
        },
        {
            "name": "microprofile-jwt",
            "description": "Microprofile - JWT built-in scope",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "true",
                "display.on.consent.screen": "false"
            },
            "protocolMappers": [
                {
                    "name": "upn",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "username",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "upn",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "groups",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-realm-role-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "multivalued": "true",
                        "user.attribute": "foo",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "groups",
                        "jsonType.label": "String"
                    }
                }
            ]
        },
        {
            "name": "roles",
            "description": "OpenID Connect scope for add user roles to the access token",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "false",
                "consent.screen.text": "${rolesScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "client roles",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-client-role-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.attribute": "foo",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "resource_access.${client_id}.roles",
                        "jsonType.label": "String",
                        "multivalued": "true"
                    }
                },
                {
                    "name": "realm roles",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-realm-role-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.attribute": "foo",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "realm_access.roles",
                        "jsonType.label": "String",
                        "multivalued": "true"
                    }
                },
                {
                    "name": "audience resolve",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-audience-resolve-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "access.token.claim": "true"
                    }
                }
            ]
        },
        {
            "name": "profile",
            "description": "OpenID Connect built-in scope: profile",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "true",
                "consent.screen.text": "${profileScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "gender",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "gender",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "gender",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "middle name",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "middleName",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "middle_name",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "locale",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "locale",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "locale",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "picture",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "picture",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "picture",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "full name",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-full-name-mapper",
                    "consentRequired": false,
                    "config": {
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "userinfo.token.claim": "true"
                    }
                },
                {
                    "name": "family name",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "lastName",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "family_name",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "profile",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "profile",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "profile",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "birthdate",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "birthdate",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "birthdate",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "website",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "website",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "website",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "nickname",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "nickname",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "nickname",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "zoneinfo",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "zoneinfo",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "zoneinfo",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "username",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "username",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "preferred_username",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "given name",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "firstName",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "given_name",
                        "jsonType.label": "String"
                    }
                },
                {
                    "name": "updated at",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "updatedAt",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "updated_at",
                        "jsonType.label": "long"
                    }
                }
            ]
        },
        {
            "name": "email",
            "description": "OpenID Connect built-in scope: email",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "true",
                "consent.screen.text": "${emailScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "email verified",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-property-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "emailVerified",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "email_verified",
                        "jsonType.label": "boolean"
                    }
                },
                {
                    "name": "email",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usermodel-attribute-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "userinfo.token.claim": "true",
                        "user.attribute": "email",
                        "id.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "email",
                        "jsonType.label": "String"
                    }
                }
            ]
        },
        {
            "name": "address",
            "description": "OpenID Connect built-in scope: address",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "true",
                "consent.screen.text": "${addressScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "address",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-address-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.attribute.formatted": "formatted",
                        "user.attribute.country": "country",
                        "introspection.token.claim": "true",
                        "user.attribute.postal_code": "postal_code",
                        "userinfo.token.claim": "true",
                        "user.attribute.street": "street",
                        "id.token.claim": "true",
                        "user.attribute.region": "region",
                        "access.token.claim": "true",
                        "user.attribute.locality": "locality"
                    }
                }
            ]
        },
        {
            "name": "basic",
            "description": "OpenID Connect scope for add all basic claims to the token",
            "protocol": "openid-connect",
            "attributes": {
                "include.in.token.scope": "false",
                "display.on.consent.screen": "false"
            },
            "protocolMappers": [
                {
                    "name": "auth_time",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-usersessionmodel-note-mapper",
                    "consentRequired": false,
                    "config": {
                        "user.session.note": "AUTH_TIME",
                        "id.token.claim": "true",
                        "introspection.token.claim": "true",
                        "access.token.claim": "true",
                        "claim.name": "auth_time",
                        "jsonType.label": "long"
                    }
                },
                {
                    "name": "sub",
                    "protocol": "openid-connect",
                    "protocolMapper": "oidc-sub-mapper",
                    "consentRequired": false,
                    "config": {
                        "introspection.token.claim": "true",
                        "access.token.claim": "true"
                    }
                }
            ]
        },
        {
            "name": "role_list",
            "description": "SAML role list",
            "protocol": "saml",
            "attributes": {
                "consent.screen.text": "${samlRoleListScopeConsentText}",
                "display.on.consent.screen": "true"
            },
            "protocolMappers": [
                {
                    "name": "role list",
                    "protocol": "saml",
                    "protocolMapper": "saml-role-list-mapper",
                    "consentRequired": false,
                    "config": {
                        "single": "false",
                        "attribute.nameformat": "Basic",
                        "attribute.name": "Role"
                    }
                }
            ]
        }
    ],
    "defaultDefaultClientScopes": [
        "role_list",
        "profile",
        "email",
        "roles",
        "web-origins",
        "acr",
        "basic"
    ],
    "defaultOptionalClientScopes": [
        "offline_access",
        "address",
        "phone",
        "microprofile-jwt"
    ],
    "browserSecurityHeaders": {
        "contentSecurityPolicyReportOnly": "",
        "xContentTypeOptions": "nosniff",
        "referrerPolicy": "no-referrer",
        "xRobotsTag": "none",
        "xFrameOptions": "SAMEORIGIN",
        "contentSecurityPolicy": "frame-src 'self'; frame-ancestors 'self'; object-src 'none';",
        "xXSSProtection": "1; mode=block",
        "strictTransportSecurity": "max-age=31536000; includeSubDomains"
    },
    "smtpServer": {},
    "eventsEnabled": false,
    "eventsListeners": [
        "jboss-logging"
    ],
    "enabledEventTypes": [],
    "adminEventsEnabled": false,
    "adminEventsDetailsEnabled": false,
    "identityProviders": [],
    "identityProviderMappers": [],
    "components": {
        "org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy": [
            {
                "name": "Consent Required",
                "providerId": "consent-required",
                "subType": "anonymous",
                "subComponents": {},
                "config": {}
            },
            {
                "name": "Allowed Client Scopes",
                "providerId": "allowed-client-templates",
                "subType": "anonymous",
                "subComponents": {},
                "config": {
                    "allow-default-scopes": [
                        "true"
                    ]
                }
            },
            {
                "name": "Full Scope Disabled",
                "providerId": "scope",
                "subType": "anonymous",
                "subComponents": {},
                "config": {}
            },
            {
                "name": "Trusted Hosts",
                "providerId": "trusted-hosts",
                "subType": "anonymous",
                "subComponents": {},
                "config": {
                    "host-sending-registration-request-must-match": [
                        "true"
                    ],
                    "client-uris-must-match": [
                        "true"
                    ]
                }
            },
            {
                "name": "Allowed Protocol Mapper Types",
                "providerId": "allowed-protocol-mappers",
                "subType": "authenticated",
                "subComponents": {},
                "config": {
                    "allowed-protocol-mapper-types": [
                        "saml-user-attribute-mapper",
                        "oidc-full-name-mapper",
                        "oidc-address-mapper",
                        "oidc-sha256-pairwise-sub-mapper",
                        "oidc-usermodel-attribute-mapper",
                        "saml-role-list-mapper",
                        "oidc-usermodel-property-mapper",
                        "saml-user-property-mapper"
                    ]
                }
            },
            {
                "name": "Allowed Client Scopes",
                "providerId": "allowed-client-templates",
                "subType": "authenticated",
                "subComponents": {},
                "config": {
                    "allow-default-scopes": [
                        "true"
                    ]
                }
            },
            {
                "name": "Allowed Protocol Mapper Types",
                "providerId": "allowed-protocol-mappers",
                "subType": "anonymous",
                "subComponents": {},
                "config": {
                    "allowed-protocol-mapper-types": [
                        "oidc-usermodel-property-mapper",
                        "saml-user-attribute-mapper",
                        "oidc-usermodel-attribute-mapper",
                        "saml-user-property-mapper",
                        "oidc-full-name-mapper",
                        "oidc-address-mapper",
                        "oidc-sha256-pairwise-sub-mapper",
                        "saml-role-list-mapper"
                    ]
                }
            },
            {
                "name": "Max Clients Limit",
                "providerId": "max-clients",
                "subType": "anonymous",
                "subComponents": {},
                "config": {
                    "max-clients": [
                        "200"
                    ]
                }
            }
        ],
        "org.keycloak.userprofile.UserProfileProvider": [
            {
                "providerId": "declarative-user-profile",
                "subComponents": {},
                "config": {
                    "kc.user.profile.config": [
                        "{\"attributes\":[{\"name\":\"username\",\"displayName\":\"${username}\",\"validations\":{\"length\":{\"min\":3,\"max\":255},\"username-prohibited-characters\":{},\"up-username-not-idn-homograph\":{}},\"permissions\":{\"view\":[\"admin\",\"user\"],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false},{\"name\":\"email\",\"displayName\":\"${email}\",\"validations\":{\"email\":{},\"length\":{\"max\":255}},\"required\":{\"roles\":[\"user\"]},\"permissions\":{\"view\":[\"admin\",\"user\"],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false},{\"name\":\"firstName\",\"displayName\":\"${firstName}\",\"validations\":{\"length\":{\"max\":255},\"person-name-prohibited-characters\":{}},\"required\":{\"roles\":[\"user\"]},\"permissions\":{\"view\":[\"admin\",\"user\"],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false},{\"name\":\"lastName\",\"displayName\":\"${lastName}\",\"validations\":{\"length\":{\"max\":255},\"person-name-prohibited-characters\":{}},\"required\":{\"roles\":[\"user\"]},\"permissions\":{\"view\":[\"admin\",\"user\"],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false},{\"name\":\"profileImage\",\"displayName\":\"Profile Image\",\"validations\":{},\"annotations\":{},\"permissions\":{\"view\":[],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false},{\"name\":\"dbid\",\"displayName\":\"Database User ID\",\"validations\":{},\"annotations\":{},\"permissions\":{\"view\":[],\"edit\":[\"admin\",\"user\"]},\"multivalued\":false}],\"groups\":[{\"name\":\"user-metadata\",\"displayHeader\":\"User metadata\",\"displayDescription\":\"Attributes, which refer to user metadata\"}]}"
                    ]
                }
            }
        ],
        "org.keycloak.keys.KeyProvider": [
            {
                "name": "rsa-generated",
                "providerId": "rsa-generated",
                "subComponents": {},
                "config": {
                    "priority": [
                        "100"
                    ]
                }
            },
            {
                "name": "hmac-generated-hs512",
                "providerId": "hmac-generated",
                "subComponents": {},
                "config": {
                    "priority": [
                        "100"
                    ],
                    "algorithm": [
                        "HS512"
                    ]
                }
            },
            {
                "name": "aes-generated",
                "providerId": "aes-generated",
                "subComponents": {},
                "config": {
                    "priority": [
                        "100"
                    ]
                }
            },
            {
                "name": "rsa-enc-generated",
                "providerId": "rsa-enc-generated",
                "subComponents": {},
                "config": {
                    "priority": [
                        "100"
                    ],
                    "algorithm": [
                        "RSA-OAEP"
                    ]
                }
            }
        ]
    },
    "internationalizationEnabled": false,
    "supportedLocales": [],
    "authenticationFlows": [
        {
            "alias": "Account verification options",
            "description": "Method with which to verity the existing account",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "idp-email-verification",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "ALTERNATIVE",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "Verify Existing Account by Re-authentication",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "Browser - Conditional OTP",
            "description": "Flow to determine if the OTP is required for the authentication",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "conditional-user-configured",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "auth-otp-form",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "Direct Grant - Conditional OTP",
            "description": "Flow to determine if the OTP is required for the authentication",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "conditional-user-configured",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "direct-grant-validate-otp",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "First broker login - Conditional OTP",
            "description": "Flow to determine if the OTP is required for the authentication",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "conditional-user-configured",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "auth-otp-form",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "Handle Existing Account",
            "description": "Handle what to do if there is existing account with same email/username like authenticated identity provider",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "idp-confirm-link",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "Account verification options",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "Reset - Conditional OTP",
            "description": "Flow to determine if the OTP should be reset or not. Set to REQUIRED to force.",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "conditional-user-configured",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "reset-otp",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "User creation or linking",
            "description": "Flow for the existing/non-existing user alternatives",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticatorConfig": "create unique user config",
                    "authenticator": "idp-create-user-if-unique",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "ALTERNATIVE",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "Handle Existing Account",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "Verify Existing Account by Re-authentication",
            "description": "Reauthentication of existing account",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "idp-username-password-form",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "CONDITIONAL",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "First broker login - Conditional OTP",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "browser",
            "description": "browser based authentication",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "auth-cookie",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "auth-spnego",
                    "authenticatorFlow": false,
                    "requirement": "DISABLED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "identity-provider-redirector",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 25,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "ALTERNATIVE",
                    "priority": 30,
                    "autheticatorFlow": true,
                    "flowAlias": "forms",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "clients",
            "description": "Base authentication for clients",
            "providerId": "client-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "client-secret",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "client-jwt",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "client-secret-jwt",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 30,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "client-x509",
                    "authenticatorFlow": false,
                    "requirement": "ALTERNATIVE",
                    "priority": 40,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "direct grant",
            "description": "OpenID Connect Resource Owner Grant",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "direct-grant-validate-username",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "direct-grant-validate-password",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "CONDITIONAL",
                    "priority": 30,
                    "autheticatorFlow": true,
                    "flowAlias": "Direct Grant - Conditional OTP",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "docker auth",
            "description": "Used by Docker clients to authenticate against the IDP",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "docker-http-basic-authenticator",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "first broker login",
            "description": "Actions taken after first broker login with identity provider account, which is not yet linked to any Keycloak account",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticatorConfig": "review profile config",
                    "authenticator": "idp-review-profile",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "User creation or linking",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "forms",
            "description": "Username, password, otp and other auth forms.",
            "providerId": "basic-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "auth-username-password-form",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "CONDITIONAL",
                    "priority": 20,
                    "autheticatorFlow": true,
                    "flowAlias": "Browser - Conditional OTP",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "registration",
            "description": "registration flow",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "registration-page-form",
                    "authenticatorFlow": true,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": true,
                    "flowAlias": "registration form",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "registration form",
            "description": "registration form",
            "providerId": "form-flow",
            "topLevel": false,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "registration-user-creation",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "registration-password-action",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 50,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "registration-recaptcha-action",
                    "authenticatorFlow": false,
                    "requirement": "DISABLED",
                    "priority": 60,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "registration-terms-and-conditions",
                    "authenticatorFlow": false,
                    "requirement": "DISABLED",
                    "priority": 70,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "reset credentials",
            "description": "Reset credentials for a user if they forgot their password or something",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "reset-credentials-choose-user",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "reset-credential-email",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 20,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticator": "reset-password",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 30,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                },
                {
                    "authenticatorFlow": true,
                    "requirement": "CONDITIONAL",
                    "priority": 40,
                    "autheticatorFlow": true,
                    "flowAlias": "Reset - Conditional OTP",
                    "userSetupAllowed": false
                }
            ]
        },
        {
            "alias": "saml ecp",
            "description": "SAML ECP Profile Authentication Flow",
            "providerId": "basic-flow",
            "topLevel": true,
            "builtIn": true,
            "authenticationExecutions": [
                {
                    "authenticator": "http-basic-authenticator",
                    "authenticatorFlow": false,
                    "requirement": "REQUIRED",
                    "priority": 10,
                    "autheticatorFlow": false,
                    "userSetupAllowed": false
                }
            ]
        }
    ],
    "authenticatorConfig": [
        {
            "alias": "create unique user config",
            "config": {
                "require.password.update.after.registration": "false"
            }
        },
        {
            "alias": "review profile config",
            "config": {
                "update.profile.on.first.login": "missing"
            }
        }
    ],
    "requiredActions": [
        {
            "alias": "CONFIGURE_TOTP",
            "name": "Configure OTP",
            "providerId": "CONFIGURE_TOTP",
            "enabled": true,
            "defaultAction": false,
            "priority": 10,
            "config": {}
        },
        {
            "alias": "TERMS_AND_CONDITIONS",
            "name": "Terms and Conditions",
            "providerId": "TERMS_AND_CONDITIONS",
            "enabled": false,
            "defaultAction": false,
            "priority": 20,
            "config": {}
        },
        {
            "alias": "UPDATE_PASSWORD",
            "name": "Update Password",
            "providerId": "UPDATE_PASSWORD",
            "enabled": true,
            "defaultAction": false,
            "priority": 30,
            "config": {}
        },
        {
            "alias": "UPDATE_PROFILE",
            "name": "Update Profile",
            "providerId": "UPDATE_PROFILE",
            "enabled": true,
            "defaultAction": false,
            "priority": 40,
            "config": {}
        },
        {
            "alias": "VERIFY_EMAIL",
            "name": "Verify Email",
            "providerId": "VERIFY_EMAIL",
            "enabled": true,
            "defaultAction": false,
            "priority": 50,
            "config": {}
        },
        {
            "alias": "delete_account",
            "name": "Delete Account",
            "providerId": "delete_account",
            "enabled": false,
            "defaultAction": false,
            "priority": 60,
            "config": {}
        },
        {
            "alias": "webauthn-register",
            "name": "Webauthn Register",
            "providerId": "webauthn-register",
            "enabled": true,
            "defaultAction": false,
            "priority": 70,
            "config": {}
        },
        {
            "alias": "webauthn-register-passwordless",
            "name": "Webauthn Register Passwordless",
            "providerId": "webauthn-register-passwordless",
            "enabled": true,
            "defaultAction": false,
            "priority": 80,
            "config": {}
        },
        {
            "alias": "VERIFY_PROFILE",
            "name": "Verify Profile",
            "providerId": "VERIFY_PROFILE",
            "enabled": true,
            "defaultAction": false,
            "priority": 90,
            "config": {}
        },
        {
            "alias": "delete_credential",
            "name": "Delete Credential",
            "providerId": "delete_credential",
            "enabled": true,
            "defaultAction": false,
            "priority": 100,
            "config": {}
        },
        {
            "alias": "update_user_locale",
            "name": "Update User Locale",
            "providerId": "update_user_locale",
            "enabled": true,
            "defaultAction": false,
            "priority": 1000,
            "config": {}
        }
    ],
    "browserFlow": "browser",
    "registrationFlow": "registration",
    "directGrantFlow": "direct grant",
    "resetCredentialsFlow": "reset credentials",
    "clientAuthenticationFlow": "clients",
    "dockerAuthenticationFlow": "docker auth",
    "firstBrokerLoginFlow": "first broker login",
    "attributes": {
        "cibaBackchannelTokenDeliveryMode": "poll",
        "cibaAuthRequestedUserHint": "login_hint",
        "oauth2DevicePollingInterval": "5",
        "clientOfflineSessionMaxLifespan": "0",
        "clientSessionIdleTimeout": "0",
        "actionTokenGeneratedByUserLifespan.verify-email": "",
        "actionTokenGeneratedByUserLifespan.idp-verify-account-via-email": "",
        "clientOfflineSessionIdleTimeout": "0",
        "actionTokenGeneratedByUserLifespan.execute-actions": "",
        "cibaInterval": "5",
        "realmReusableOtpCode": "false",
        "cibaExpiresIn": "120",
        "oauth2DeviceCodeLifespan": "600",
        "parRequestUriLifespan": "60",
        "clientSessionMaxLifespan": "0",
        "organizationsEnabled": "false",
        "shortVerificationUri": "",
        "actionTokenGeneratedByUserLifespan.reset-credentials": ""
    },
    "keycloakVersion": "25.0.0",
    "userManagedAccessAllowed": false,
    "organizationsEnabled": false,
    "clientProfiles": {
        "profiles": []
    },
    "clientPolicies": {
        "policies": []
    }
}
`
