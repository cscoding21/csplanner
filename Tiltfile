load('./tilt_lib/csserver/Tiltfile', 'setup_csserver')
load('./tilt_lib/debug/Tiltfile', 'setup_debug')
load('./tilt_lib/surrealdb/Tiltfile', 'setup_surrealdb')
load('./tilt_lib/ollama/Tiltfile', 'setup_ollama')


# DEPLOY SERVICES
#---- Third party
setup_surrealdb()
setup_ollama()

#---- Tools
setup_debug()

#---- App
setup_csserver() 



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
  labels=["Analyzer"],
  resource_deps=['surrealdb']
)
