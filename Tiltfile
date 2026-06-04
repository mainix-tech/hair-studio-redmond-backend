# Load the restart_process extension
load('ext://restart_process', 'docker_build_with_restart')

### K8s Config ###

# Uncomment to use secrets
# k8s_yaml('./infra/development/k8s/secrets.yaml')

k8s_yaml('./infra/development/k8s/app-config.yaml')

### End of K8s Config ###
### API Gateway ###

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway'
if os.name == 'nt':
  gateway_compile_cmd = './infra/development/docker/api-gateway-build.bat'

local_resource(
  'api-gateway-compile',
  gateway_compile_cmd,
  deps=['./services/api-gateway', './shared'],
  labels="compiles",
)

docker_build_with_restart(
  'hair-studio-redmond/api-gateway',
  '.',
  entrypoint=['/app/build/api-gateway'],
  dockerfile='./infra/development/docker/api-gateway.Dockerfile',
  only=[
    './build/api-gateway',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/api-gateway-deployment.yaml')
k8s_resource(
  'api-gateway',
  port_forwards=8081,
  resource_deps=['api-gateway-compile'],
  labels="services",
)

### End of API Gateway ###
### Profile Service ###

profile_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/profile-service ./services/profile-service/cmd/main.go'
if os.name == 'nt':
  profile_compile_cmd = './infra/development/docker/profile-service-build.bat'

local_resource(
  'profile-service-compile',
  profile_compile_cmd,
  deps=['./services/profile-service', './shared'],
  labels="compiles",
)

docker_build_with_restart(
  'hair-studio-redmond/profile-service',
  '.',
  entrypoint=['/app/build/profile-service'],
  dockerfile='./infra/development/docker/profile-service.Dockerfile',
  only=[
    './build/profile-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/profile-service-deployment.yaml')
k8s_resource(
  'profile-service',
  resource_deps=['profile-service-compile'],
  labels="services",
)

### End of Profile Service ###
### Postgres Database ###

# Apply the postgres kubernetes resources (Deployment, Service, and Initializer script)
k8s_yaml('./infra/development/k8s/postgres-db-deployment.yaml')

# Register it in Tilt so you can track database logs and availability status
k8s_resource(
  'postgres-db',
  # Port forward so you can optionally connect via Datagrip/DBeaver on your local machine
  port_forwards='5433:5432',
  labels="storage",
)

### End of Postgres Database ###