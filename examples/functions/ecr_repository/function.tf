# arn:aws:ecr:ap-northeast-1:111111111111:repository/repository-name
output "ecr_repository" {
  value = provider::arn::ecr_repository("repository-name")
}
