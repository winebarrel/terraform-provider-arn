# arn:aws:ecs:ap-northeast-1:111111111111:cluster/cluster-name
output "ecs_cluster" {
  value = provider::arn::ecs_cluster("cluster-name")
}
