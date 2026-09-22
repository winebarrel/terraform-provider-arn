# arn:aws:detective:ap-northeast-1:111111111111:graph:resource-id
output "detective_graph" {
  value = provider::arn::detective_graph("resource-id")
}
