# arn:aws:nimble:ap-northeast-1:111111111111:eula-acceptance/eula-acceptance-id
output "nimble_eula_acceptance" {
  value = provider::arn::nimble_eula_acceptance("eula-acceptance-id")
}
