# arn:aws:iotsitewise:ap-northeast-1:111111111111:project/project-id
output "iotsitewise_project" {
  value = provider::arn::iotsitewise_project("project-id")
}
