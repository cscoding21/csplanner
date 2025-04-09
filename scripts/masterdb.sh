#!/bin/bash

HAS=$(psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' -c "SELECT 1 FROM pg_database WHERE datname = 'cssaas'")

if [[ $HAS == *"1 row)" ]]; then
    echo "cssaas DB already exists"
else
    echo "Creating cssaas DB"
    psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
    CREATE DATABASE cssaas;

    CREATE USER cssaas WITH PASSWORD 'cssaas';
    GRANT ALL PRIVILEGES ON DATABASE cssaas TO cssaas;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO cssaas;
EOF
fi

psql 'dbname=cssaas host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
CREATE TABLE IF NOT EXISTS organization (
    id VARCHAR(128) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    url VARCHAR(128) NOT NULL,
    is_provisioned BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(128),
    data JSONB
);

CREATE INDEX IF NOT EXISTS idx_organization_id 
ON organization(id);

CREATE INDEX IF NOT EXISTS idx_organization_url 
ON organization(url);

CREATE TABLE IF NOT EXISTS license (
    id VARCHAR(128) UNIQUE NOT NULL,
    org_id VARCHAR(128) UNIQUE NOT NULL  REFERENCES organization (id),
    app VARCHAR(128) UNIQUE NOT NULL,
    email VARCHAR(128) NOT NULL,
    expires_on TIMESTAMP NOT NULL,
    cancelled_on TIMESTAMP NOT NULL,
    auto_renew BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(128),
    data JSONB
);

CREATE INDEX IF NOT EXISTS idx_license_id 
ON license(id);

CREATE INDEX IF NOT EXISTS idx_license_org_id 
ON license(org_id);

CREATE INDEX IF NOT EXISTS idx_license_email 
ON license(email);



EOF

