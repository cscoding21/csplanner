load('./tilt_lib/csserver/Tiltfile', 'setup_csserver')
load('./tilt_lib/csweb/Tiltfile', 'setup_csweb')
load('./tilt_lib/csai/Tiltfile', 'setup_csai')
load('./tilt_lib/csmcp/Tiltfile', 'setup_csservermcp')
load('./tilt_lib/csprocessor/Tiltfile', 'setup_csprocessor')
load('./tilt_lib/cssaas/Tiltfile', 'setup_cssaas')
load('./tilt_lib/debug/Tiltfile', 'setup_debug')
# load('./tilt_lib/surrealdb/Tiltfile', 'setup_surrealdb')
load('./tilt_lib/postgres/Tiltfile', 'setup_postgres')
load('./tilt_lib/nats/Tiltfile', 'setup_nats')
load('./tilt_lib/ollama/Tiltfile', 'setup_ollama')
load('./tilt_lib/keycloak/Tiltfile', 'setup_keycloak')
load('./tilt_lib/vault/Tiltfile', 'setup_vault')


# DEPLOY SERVICES
#---- Third party
#setup_surrealdb()
setup_postgres()
#setup_ollama()
setup_keycloak()
setup_nats()
#setup_vault()

#---- Tools
setup_debug()

#---- App
setup_csserver() 
setup_csweb() 
setup_csai()
setup_cssaas() 
setup_csservermcp()
setup_csprocessor()


###############################################################
# Watch: tell Tilt how to connect locally (optional)
# Port forward service
# SurrealDB
# k8s_resource(
#   workload='surrealdb',
#   port_forwards=9999,
#   labels=["Database"]
# )

k8s_resource(
  workload='csserver',
  port_forwards="5000:5000",
  labels=["csPlanner"] 
)
# resource_deps=['surrealdb']

k8s_resource(
  workload='csmcp',
  port_forwards="6002:6002",
  labels=["csPlanner"] 
)

k8s_resource(
  workload='csprocessor',
  labels=["csPlanner"] 
)

k8s_resource(
  workload='csweb',
  port_forwards="4000:4000",
  labels=["csPlanner"],
  resource_deps=['csserver']
)

k8s_resource(
  workload='cssaas',
  port_forwards="3006:3006",
  labels=["csPlanner"],
  resource_deps=['csserver']
)

k8s_resource(
  workload='csai',
  port_forwards="7000:7000",
  labels=["csPlanner"]
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

# k8s_resource(
#   workload='nats',
#   port_forwards="4222:4222",
#   labels=["NATS"],
# )

# local_resource(
#   'database_setup',
#   cmd='sh -c "tilt_lib/postgres/setup.sh"',
#   resource_deps=['postgres']
# )
