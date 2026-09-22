# arn:aws:ecs:ap-northeast-1:111111111111:daemon-deployment/cluster-name/daemon-name/daemon-deployment-id
output "ecs_daemon_deployment" {
  value = provider::arn::ecs_daemon_deployment("cluster-name", "daemon-name", "daemon-deployment-id")
}
