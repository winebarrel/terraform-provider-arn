# arn:aws:aps:ap-northeast-1:111111111111:workspace/workspace-id
output "aps_workspace" {
  value = provider::arn::aps_workspace("workspace-id")
}
