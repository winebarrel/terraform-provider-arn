# arn:aws:codeartifact:ap-northeast-1:111111111111:repository/domain-name/repository-name
output "codeartifact_repository" {
  value = provider::arn::codeartifact_repository("domain-name", "repository-name")
}
