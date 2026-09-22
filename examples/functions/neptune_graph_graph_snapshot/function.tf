# arn:aws:neptune-graph:ap-northeast-1:111111111111:graph-snapshot/resource-id
output "neptune_graph_graph_snapshot" {
  value = provider::arn::neptune_graph_graph_snapshot("resource-id")
}
