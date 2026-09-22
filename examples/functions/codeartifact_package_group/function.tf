# arn:aws:codeartifact:ap-northeast-1:111111111111:package-group/domain-nameencoded-package-group-pattern
output "codeartifact_package_group" {
  value = provider::arn::codeartifact_package_group("domain-name", "encoded-package-group-pattern")
}
