# arn:aws:fms:ap-northeast-1:111111111111:resource-set/id
output "fms_resource_set" {
  value = provider::arn::fms_resource_set("id")
}
