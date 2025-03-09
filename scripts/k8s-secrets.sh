#!/bin/bash

#---create app namespace if it doesn't exist
kubectl create namespace app || true

#---delete existing secrets
kubectl delete secret -n app kc-client-secret
kubectl delete secret -n app kc-admin-pass
kubectl delete secret -n app db-pass

#---recreate
kubectl create secret generic -n app kc-client-secret --from-file=./secrets/kc-client-secret.txt
kubectl create secret generic -n app kc-admin-pass --from-file=./secrets/kc-admin-pass.txt
kubectl create secret generic -n app db-pass --from-file=./secrets/db-pass.txt