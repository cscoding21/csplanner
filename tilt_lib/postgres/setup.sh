#!/bin/bash

HAS=$(psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' -c "SELECT 1 FROM pg_database WHERE datname = 'csplanner'")

if [[ $HAS == *"1 row)" ]]; then
    echo "csplanner DB already exists"
else
    echo "Creating csplanner DB"
    psql 'dbname=postgres host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
    CREATE DATABASE csplanner;

    CREATE USER csplanner WITH PASSWORD 'csplanner';
    GRANT ALL PRIVILEGES ON DATABASE csplanner TO csplanner;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO csplanner;
EOF
fi


## declare an array variable
declare -a arr=("list" "organization" "project" "projecttemplate" "resource" "role" "appuser")

## now loop through the above array
for i in "${arr[@]}"
do
    psql 'dbname=csplanner host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF
    CREATE TABLE IF NOT EXISTS ${i} (
        id VARCHAR(128) UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(128) NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        updated_by VARCHAR(128) NOT NULL,
        deleted_at TIMESTAMP,
        deleted_by VARCHAR(128),
        data JSONB
    );

    CREATE INDEX IF NOT EXISTS idx_${i}_id 
    ON ${i}(id);

    CREATE INDEX IF NOT EXISTS idx_${i}_created_at 
    ON ${i}(created_at);

    CREATE INDEX IF NOT EXISTS idx_${i}_created_by
    ON ${i}(created_by);

    CREATE INDEX IF NOT EXISTS idx_${i}_updated_at
    ON ${i}(updated_at);

    CREATE INDEX IF NOT EXISTS idx_${i}_updated_by
    ON ${i}(updated_by);
EOF
done



