# arn:aws:codebuild:ap-northeast-1:111111111111:project/project-name
output "codebuild_project" {
  value = provider::arn::codebuild_project("project-name")
}
