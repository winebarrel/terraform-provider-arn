# arn:aws:ecs:ap-northeast-1:111111111111:container-instance/cluster-name/container-instance-id
output "ecs_container_instance" {
  value = provider::arn::ecs_container_instance("cluster-name", "container-instance-id")
}
