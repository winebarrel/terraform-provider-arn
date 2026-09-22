# arn:aws:codewhisperer:ap-northeast-1:111111111111:profile/identifier
output "q_profile" {
  value = provider::arn::q_profile("identifier")
}
