# arn:aws:ecs:ap-northeast-1:111111111111:service-deployment/cluster-name/service-name/service-deployment-id
output "ecs_service_deployment" {
  value = provider::arn::ecs_service_deployment("cluster-name", "service-name", "service-deployment-id")
}
