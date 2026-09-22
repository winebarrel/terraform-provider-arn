# arn:aws:codewhisperer:ap-northeast-1:111111111111:customization/identifier
output "codewhisperer_customization" {
  value = provider::arn::codewhisperer_customization("identifier")
}
