# arn:aws:codeconnections:ap-northeast-1:111111111111:repository-link/repository-link-id
output "codeconnections_repository_link" {
  value = provider::arn::codeconnections_repository_link("repository-link-id")
}
