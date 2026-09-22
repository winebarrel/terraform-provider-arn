# arn:aws:ecs:ap-northeast-1:111111111111:daemon/cluster-name/daemon-name
output "ecs_daemon" {
  value = provider::arn::ecs_daemon("cluster-name", "daemon-name")
}
