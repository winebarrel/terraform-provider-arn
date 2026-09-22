# arn:aws:codeartifact:ap-northeast-1:111111111111:package/domain-name/repository-name/package-format/package-namespace/package-name
output "codeartifact_package" {
  value = provider::arn::codeartifact_package("domain-name", "repository-name", "package-format", "package-namespace", "package-name")
}
