# arn:aws:ecs:ap-northeast-1:111111111111:daemon-revision/cluster-name/daemon-name/daemon-revision-id
output "ecs_daemon_revision" {
  value = provider::arn::ecs_daemon_revision("cluster-name", "daemon-name", "daemon-revision-id")
}
