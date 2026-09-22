# arn:aws:nimble:ap-northeast-1:111111111111:eula/eula-id
output "nimble_eula" {
  value = provider::arn::nimble_eula("eula-id")
}
