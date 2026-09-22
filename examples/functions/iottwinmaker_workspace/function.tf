# arn:aws:iottwinmaker:ap-northeast-1:111111111111:workspace/workspace-id
output "iottwinmaker_workspace" {
  value = provider::arn::iottwinmaker_workspace("workspace-id")
}
