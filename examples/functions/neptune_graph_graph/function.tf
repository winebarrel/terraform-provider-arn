# arn:aws:neptune-graph:ap-northeast-1:111111111111:graph/resource-id
output "neptune_graph_graph" {
  value = provider::arn::neptune_graph_graph("resource-id")
}
