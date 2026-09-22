# arn:aws:inspector2:ap-northeast-1:111111111111:codesecurity-integration/code-security-integration-id
output "inspector2_code_security_integration" {
  value = provider::arn::inspector2_code_security_integration("code-security-integration-id")
}
