# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/workspace/workspace-id
output "connect_workspace" {
  value = provider::arn::connect_workspace("instance-id", "workspace-id")
}
