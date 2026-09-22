# arn:aws:lookoutvision:ap-northeast-1:111111111111:project/project-name
output "lookoutvision_project" {
  value = provider::arn::lookoutvision_project("project-name")
}
