# arn:aws:evidently:ap-northeast-1:111111111111:project/project-name
output "evidently_project" {
  value = provider::arn::evidently_project("project-name")
}
