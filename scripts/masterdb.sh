#!/bin/bash

HAS=$(psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' -c "SELECT 1 FROM pg_database WHERE datname = 'cssaas'")

if [[ $HAS == *"1 row)" ]]; then
    echo "cssaas DB already exists"
else
    echo "Creating cssaas DB"
    psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
    CREATE DATABASE cssaas;

    CREATE USER usr_cssaas WITH PASSWORD 'cssaas';
    GRANT ALL PRIVILEGES ON DATABASE cssaas TO usr_cssaas;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO usr_cssaas;
EOF
fi

psql 'dbname=cssaas host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
CREATE TABLE IF NOT EXISTS organization (
    id VARCHAR(128) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    db_host VARCHAR(128) NOT NULL,
    database VARCHAR(128) NOT NULL,
    url_key VARCHAR(128) UNIQUE NOT NULL,
    realm VARCHAR(128) NOT NULL,
    is_provisioned BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(128) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(128),
    data JSONB
);

CREATE INDEX IF NOT EXISTS idx_organization_id 
ON organization(id);

CREATE INDEX IF NOT EXISTS idx_organization_url_key 
ON organization(url_key);

CREATE TABLE IF NOT EXISTS license (
    id VARCHAR(128) UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    org_id VARCHAR(128) UNIQUE NOT NULL  REFERENCES organization (id),
    app VARCHAR(128) UNIQUE NOT NULL,
    email VARCHAR(128) NOT NULL,
    expires_on TIMESTAMP NOT NULL,
    cancelled_on TIMESTAMP,
    auto_renew BOOLEAN NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(128) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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


psql 'dbname=cssaas host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF

INSERT INTO organization (
    id, 
    name, 
    db_host, 
    database, 
    url_key, 
    realm, 
    is_provisioned, 
    created_at, 
    created_by, 
    updated_at, 
    updated_by
) VALUES (
    'organization:jeph', 
    'Jeph Test Org', 
    'localhost', 
    'csp_jeph_test_org', 
    'localhost', 
    'jeph_test_org', 
    false, 
    CURRENT_TIMESTAMP, 
    'provisioner', 
    CURRENT_TIMESTAMP, 
    'provisioner'
);

INSERT INTO license (
    org_id, 
    app, 
    email, 
    expires_on, 
    cancelled_on, 
    auto_renew, 
    created_at, 
    created_by, 
    updated_at, 
    updated_by
) VALUES (
    'organization:jeph', 
    'csplanner', 
    'jeph@jmk21.com', 
    (now() + interval '30 days'), 
    NULL, 
    true, 
    CURRENT_TIMESTAMP, 
    'provisioner', 
    CURRENT_TIMESTAMP, 
    'provisioner'
);

EOF