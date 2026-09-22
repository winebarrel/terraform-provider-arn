# arn:aws:medialive:ap-northeast-1:111111111111:cluster:cluster-id
output "medialive_cluster" {
  value = provider::arn::medialive_cluster("cluster-id")
}
