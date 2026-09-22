# arn:aws:codestar-connections:ap-northeast-1:111111111111:repository-link/repository-link-id
output "codestar_connections_repository_link" {
  value = provider::arn::codestar_connections_repository_link("repository-link-id")
}
