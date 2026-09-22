# arn:aws:medialive:ap-northeast-1:111111111111:node:cluster-id/node-id
output "medialive_node" {
  value = provider::arn::medialive_node("cluster-id", "node-id")
}
