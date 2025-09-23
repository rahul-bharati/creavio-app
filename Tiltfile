env = os.getenv("ENV", "dev")

docker_build("creavio-frontend-dev", "./web", dockerfile="./web/dev.Dockerfile", live_update=[
    sync("./web", "/app")
])
k8s_yaml("infra/k8s/dev/web-deployment.yaml")
k8s_resource("creavio-frontend-dev", port_forwards=["3000:3000"])
