# Helm Repos

- helm repo add bitnami https://charts.bitnami.com/bitnami
- helm repo add nats https://nats-io.github.io/k8s/helm/charts/
- helm repo add codecentric https://codecentric.github.io/helm-charts


# NATS
- docker run -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats-server -ti nats:latest -js

CLI Commands
- nats stream purge csplanner