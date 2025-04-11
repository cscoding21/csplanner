package provision

var checkDatabaseExistSQL = `
	SELECT 1 FROM pg_database WHERE datname = $1
`

var newDatabaseSQL = "CREATE DATABASE %s;"
var newDBUserSQL = "CREATE USER %s WITH PASSWORD '%s';"
var grantOnDBSQL = "GRANT ALL PRIVILEGES ON DATABASE %s TO %s;"
var grandOnSchemaSQL = "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO %s;"

var newMasterOrgRecordSQL = `
INSERT INTO organization (
	id,
	name,
	url,
	is_provisioned,
    db_host,
	created_by,
	updated_by
	) VALUES (
	 $1, $2, $3, $4, $5, $6, $7
);
`

var newLicenseSQL = `
INSERT INTO license (
	id,
	org_id,
	app,
	email,
	expires_on
	cancelled_on,
	auto_renew,
	created_by,
	updated_by
	) VALUES (
	 $1, $2, $3, $4, $5, $6, $7, $8, $9
);
`

var setProvisionedCompleteSQL = `
UPDATE organization
SET is_provisioned = true,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $1
WHERE id = $2;
`

var createCSPlannerTableSQL = `
CREATE TABLE IF NOT EXISTS %s (
        id VARCHAR(128) UNIQUE NOT NULL,
        parent_id VARCHAR(128),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(128) NOT NULL,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_by VARCHAR(128) NOT NULL,
        deleted_at TIMESTAMP ,
        deleted_by VARCHAR(128),
        data JSONB
    );

    CREATE INDEX IF NOT EXISTS idx_%s_id 
    ON %s(id);

    CREATE INDEX IF NOT EXISTS idx_%s_parent_id 
    ON %s(parent_id);

    CREATE INDEX IF NOT EXISTS idx_%s_created_at 
    ON %s(created_at);

    CREATE INDEX IF NOT EXISTS idx_%s_created_by
    ON %s(created_by);

    CREATE INDEX IF NOT EXISTS idx_%s_updated_at
    ON %s(updated_at);

    CREATE INDEX IF NOT EXISTS idx_%s_updated_by
    ON %s(updated_by);
`
