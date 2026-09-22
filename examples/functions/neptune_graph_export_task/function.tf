# arn:aws:neptune-graph:ap-northeast-1:111111111111:export-task/resource-id
output "neptune_graph_export_task" {
  value = provider::arn::neptune_graph_export_task("resource-id")
}
