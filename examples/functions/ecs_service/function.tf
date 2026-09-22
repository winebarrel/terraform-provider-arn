# arn:aws:ecs:ap-northeast-1:111111111111:service/cluster-name/service-name
output "ecs_service" {
  value = provider::arn::ecs_service("cluster-name", "service-name")
}
