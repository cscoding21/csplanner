package provision

import (
	"context"
	"csserver/internal/config"

	"github.com/Nerzal/gocloak/v13"

	log "github.com/sirupsen/logrus"
)

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
