#!/bin/bash

if [[ $(minikube status | grep Running) ]]; then
    echo "Minikube already running"
else
    echo "Minikube not running: starting..."

    HAS_STATE=$(minikube status | grep Stopped)
    echo "HAS_STATE = ${HAS_STATE}"
    echo "\n\n"
    

    if [[ -z "${HAS_STATE}" ]]; then
        #---configure minikube
        echo "Setting minikube resources"
        minikube config set cpus 8  
        minikube config set memory 16384
    fi

    #---start minikube
    # --gpus all 
    minikube start --kubernetes-version=v1.31.0 --driver=docker --container-runtime docker --mount-string="${HOME}/projects/data:/mnt/data" --mount

    kubectl apply -f scripts/k8s/sc.yaml

    if [[ -z ${HAS_STATE} ]]; then
        echo "Setting ingress and linkerd"

        #---enable minikube addons
        minikube addons enable ingress
        minikube addons enable nvidia-device-plugin

        #---below are only supported with KVM driver
        # minikube addons enable nvidia-gpu-device-plugin
        # minikube addons enable nvidia-driver-installer

        #---check for linkerd dependencies...for initial setup
        linkerd check --pre

        #---install linkerd crds
        linkerd install --crds | kubectl apply -f -

        #---install linkerd
        linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -

        #---install linkerd dashboard
        linkerd viz install | kubectl apply -f -
    fi
fi

./scripts/k8s-secrets.sh

echo "Running tilt up"
tilt up
