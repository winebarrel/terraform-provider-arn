# arn:aws:codestar:ap-northeast-1:111111111111:project/project-id
output "codestar_project" {
  value = provider::arn::codestar_project("project-id")
}
