env = os.getenv("ENV", "dev")

# Common settings
k8s_yaml(local('kubectl kustomize infra/k8s/dev'))
watch_file("infra/k8s/dev/config.env")

# Frontend service
docker_build("creavio-frontend-dev", "./web", dockerfile="./web/dev.Dockerfile", live_update=[
    sync("./web", "/app")
])
# k8s_yaml("infra/k8s/dev/web-deployment.yaml") # Uncomment if you want to manage k8s resources separately

# Backend service - API GATEWAY
docker_build("creavio-api-gateway-dev", ".", dockerfile="./services/api-gateway/dev.Dockerfile", live_update=[
    sync("./services/api-gateway", "/src/services/api-gateway"),
])

# Backend service - USER SERVICE
docker_build("creavio-user-service-dev", ".", dockerfile="./services/user/dev.Dockerfile", live_update=[
    sync("./services/user", "/src/services/user"),
])

# Backend service - NOTIFICATION SERVICE
docker_build("creavio-notification-service-dev", ".", dockerfile="./services/notification/dev.Dockerfile", live_update=[
    sync("./services/notification", "/src/services/notification"),
])

# --- Nice port-forwards (resource names from your Deployments)
k8s_resource('creavio-frontend-dev',            port_forwards=['3000:3000'])
k8s_resource('creavio-api-gateway-dev',         port_forwards=['8000:8000'])
k8s_resource('creavio-user-service-dev',        port_forwards=['8001:8001'])
k8s_resource('creavio-notification-service-dev',port_forwards=['8002:8002'])