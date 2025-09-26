env = os.getenv("ENV", "dev")

# Frontend service

docker_build("creavio-frontend-dev", "./web", dockerfile="./web/dev.Dockerfile", live_update=[
    sync("./web", "/app")
])
k8s_yaml("infra/k8s/dev/web-deployment.yaml")
k8s_resource("creavio-frontend-dev", port_forwards=["3000:3000"])

# Backend service - API GATEWAY
docker_build("creavio-api-gateway-dev", ".", dockerfile="./services/api-gateway/dev.Dockerfile", live_update=[
    sync("./services/api-gateway", "/src/services/api-gateway"),
])
k8s_yaml("infra/k8s/dev/api-gateway-deployment.yaml")
k8s_resource("creavio-api-gateway-dev", port_forwards=["8000:8000"])

# Backend service - USER SERVICE
docker_build("creavio-user-service-dev", ".", dockerfile="./services/user/dev.Dockerfile", live_update=[
    sync("./services/user", "/src/services/user"),
])
k8s_yaml("infra/k8s/dev/user-service-deployment.yaml")
k8s_resource("creavio-user-service-dev", port_forwards=["8001:8001"])

# Backend service - NOTIFICATION SERVICE
docker_build("creavio-notification-service-dev", ".", dockerfile="./services/notification/dev.Dockerfile", live_update=[
    sync("./services/notification", "/src/services/notification"),
])
k8s_yaml("infra/k8s/dev/notification-service-deployment.yaml")
k8s_resource("creavio-notification-service-dev", port_forwards=["8002:8002"])