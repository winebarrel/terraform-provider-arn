# arn:aws:ivs:ap-northeast-1:111111111111:composition/resource-id
output "ivs_composition" {
  value = provider::arn::ivs_composition("resource-id")
}
