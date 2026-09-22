# arn:aws:neptune-graph:ap-northeast-1:111111111111:import-task/resource-id
output "neptune_graph_import_task" {
  value = provider::arn::neptune_graph_import_task("resource-id")
}
