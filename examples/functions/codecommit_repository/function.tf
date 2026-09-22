# arn:aws:codecommit:ap-northeast-1:111111111111:repository-name
output "codecommit_repository" {
  value = provider::arn::codecommit_repository("repository-name")
}
