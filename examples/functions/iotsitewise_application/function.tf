# arn:aws:iotsitewise:ap-northeast-1:111111111111:workspace/workspace-name/application/application-id
output "iotsitewise_application" {
  value = provider::arn::iotsitewise_application("workspace-name", "application-id")
}
