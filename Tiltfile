load('./tilt_lib/csserver/Tiltfile', 'setup_csserver')
load('./tilt_lib/csweb/Tiltfile', 'setup_csweb')
load('./tilt_lib/debug/Tiltfile', 'setup_debug')
load('./tilt_lib/surrealdb/Tiltfile', 'setup_surrealdb')
load('./tilt_lib/postgres/Tiltfile', 'setup_postgres')
load('./tilt_lib/nats/Tiltfile', 'setup_nats')
load('./tilt_lib/ollama/Tiltfile', 'setup_ollama')
load('./tilt_lib/keycloak/Tiltfile', 'setup_keycloak')


# DEPLOY SERVICES
#---- Third party
setup_surrealdb()
setup_postgres()
# setup_ollama()
setup_keycloak()
setup_nats()

#---- Tools
setup_debug()

#---- App
setup_csserver() 
setup_csweb() 



###############################################################
# Watch: tell Tilt how to connect locally (optional)
# Port forward service
# SurrealDB
k8s_resource(
  workload='surrealdb',
  port_forwards=9999,
  labels=["Database"]
)

k8s_resource(
  workload='csserver',
  port_forwards="5000:5000",
  labels=["csPlanner"],
  resource_deps=['surrealdb']
)

k8s_resource(
  workload='csweb',
  port_forwards="4000:4000",
  labels=["csPlanner"],
  resource_deps=['csserver']
)

k8s_resource(
  workload='postgres',
  port_forwards="5432:5432",
  labels=["Database"],
)

k8s_resource(
  workload='keycloak',
  port_forwards="8080:8080",
  labels=["Keycloak"]
)
