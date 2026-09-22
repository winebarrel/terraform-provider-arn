# arn:aws:ecs:ap-northeast-1:111111111111:service-revision/cluster-name/service-name/service-revision-id
output "ecs_service_revision" {
  value = provider::arn::ecs_service_revision("cluster-name", "service-name", "service-revision-id")
}
