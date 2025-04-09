#!/bin/bash

## declare an array variable or tables to create
declare -a arr=("list" "organization" "project" "projecttemplate" "resource" "role" "appuser" "comment" "reaction")

## now loop through the above array
for i in "${arr[@]}"
do
    psql 'dbname=csplanner host=localhost user=postgres password=postgres port=5432 sslmode=disable' << EOF

    DELETE FROM ${i};
EOF
done