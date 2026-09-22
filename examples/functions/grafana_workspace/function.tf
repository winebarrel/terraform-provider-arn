# arn:aws:grafana:ap-northeast-1:111111111111:/workspaces/resource-id
output "grafana_workspace" {
  value = provider::arn::grafana_workspace("resource-id")
}
