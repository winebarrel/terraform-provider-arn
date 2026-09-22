# arn:aws:batch:ap-northeast-1:111111111111:consumable-resource/consumable-resource-name
output "batch_consumable_resource" {
  value = provider::arn::batch_consumable_resource("consumable-resource-name")
}
