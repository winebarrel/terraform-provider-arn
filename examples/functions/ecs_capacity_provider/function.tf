# arn:aws:ecs:ap-northeast-1:111111111111:capacity-provider/capacity-provider-name
output "ecs_capacity_provider" {
  value = provider::arn::ecs_capacity_provider("capacity-provider-name")
}
