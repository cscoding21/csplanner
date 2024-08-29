#!/bin/bash

kubectl create secret generic -n app kc-client-secret --from-file=./secrets/kc-client-secret.txt

kubectl create secret generic -n app kc-admin-pass --from-file=./secrets/kc-admin-pass.txt

kubectl create secret generic -n app surreal-pass --from-file=./secrets/surreal-pass.txt