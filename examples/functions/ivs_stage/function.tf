# arn:aws:ivs:ap-northeast-1:111111111111:stage/resource-id
output "ivs_stage" {
  value = provider::arn::ivs_stage("resource-id")
}
