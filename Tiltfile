# Load the restart_process extension
load('ext://restart_process', 'docker_build_with_restart')

### K8s Config ###

# Uncomment to use secrets
# k8s_yaml('./infra/development/k8s/secrets.yaml')

k8s_yaml('./infra/development/k8s/app-config.yaml')

### End of K8s Config ###
### API Gateway ###

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./service/api-gateway'
if os.name == 'nt':
  gateway_compile_cmd = './infra/development/docker/api-gateway-build.bat'

local_resource(
  'api-gateway-compile',
  gateway_compile_cmd,
  deps=['./service/api-gateway', './shared'],
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
  labels="service",
)

### End of API Gateway ###
### Profile Service ###

profile_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/profile-service ./service/profile-service/cmd/main.go'
if os.name == 'nt':
  profile_compile_cmd = './infra/development/docker/profile-service-build.bat'

local_resource(
  'profile-service-compile',
  profile_compile_cmd,
  deps=['./service/profile-service', './shared'],
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
  labels="service",
)

### End of Profile Service ###

### Menu Service ###

menu_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/menu-service ./service/menu-service/cmd/main.go'
if os.name == 'nt':
  menu_compile_cmd = './infra/development/docker/menu-service-build.bat'

local_resource(
  'menu-service-compile',
  menu_compile_cmd,
  deps=['./service/menu-service', './shared'],
  labels="compiles",
)

docker_build_with_restart(
  'hair-studio-redmond/menu-service',
  '.',
  entrypoint=['/app/build/menu-service'],
  dockerfile='./infra/development/docker/menu-service.Dockerfile',
  only=[
    './build/menu-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/menu-service-deployment.yaml')
k8s_resource(
  'menu-service',
  resource_deps=['menu-service-compile'],
  labels="service",
)

### End of Menu Service ###

### Catalog Service ###
catalog_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/catalog-service ./service/catalog-service/cmd/main.go'
if os.name == 'nt':
  catalog_compile_cmd = './infra/development/docker/catalog-service-build.bat'

local_resource(
  'catalog-service-compile',
  catalog_compile_cmd,
  deps=['./service/catalog-service', './shared'],
  labels="compiles",
)

docker_build_with_restart(
  'hair-studio-redmond/catalog-service',
  '.',
  entrypoint=['/app/build/catalog-service'],
  dockerfile='./infra/development/docker/catalog-service.Dockerfile',
  only=[
    './build/catalog-service',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/catalog-service-deployment.yaml')
k8s_resource(
  'catalog-service',
  resource_deps=['catalog-service-compile'],
  labels="service",
)

### End of Catalog Service ###


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