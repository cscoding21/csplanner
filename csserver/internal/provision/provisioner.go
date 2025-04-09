package provision

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ProvisionNewOrganization(
	ctx context.Context,
	db *pgxpool.Pool,
	name string) {

	/*
		check for existing name and provision status in master DB table
		- check for existing db
			-  if yes, exit

		create db for new org

		run setup scripts/migrations
		- create bot user
		- create default organization
		- create initial lists
			- skills (no values)
			- funding sources (internal/external)
			- value categories ()
		- create project templates
			- simple
			- complex

	*/
}
