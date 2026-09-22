# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/attributes/attribute-type
output "mobiletargeting_attribute" {
  value = provider::arn::mobiletargeting_attribute("app-id", "attribute-type")
}
