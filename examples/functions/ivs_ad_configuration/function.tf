# arn:aws:ivs:ap-northeast-1:111111111111:ad-configuration/resource-id
output "ivs_ad_configuration" {
  value = provider::arn::ivs_ad_configuration("resource-id")
}
