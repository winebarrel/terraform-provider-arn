# arn:aws:codewhisperer:ap-northeast-1:111111111111:profile/identifier
output "codewhisperer_profile" {
  value = provider::arn::codewhisperer_profile("identifier")
}
