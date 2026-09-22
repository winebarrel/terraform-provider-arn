# arn:aws:ecr-public::111111111111:repository/repository-name
output "ecr_public_repository" {
  value = provider::arn::ecr_public_repository("repository-name")
}
