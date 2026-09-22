# arn:aws:iotsitewise:ap-northeast-1:111111111111:workspace/workspace-name
output "iotsitewise_workspace" {
  value = provider::arn::iotsitewise_workspace("workspace-name")
}
