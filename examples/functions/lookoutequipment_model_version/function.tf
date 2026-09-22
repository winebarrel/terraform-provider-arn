# arn:aws:lookoutequipment:ap-northeast-1:111111111111:model/model-name/model-id/model-version/model-version-number
output "lookoutequipment_model_version" {
  value = provider::arn::lookoutequipment_model_version("model-name", "model-id", "model-version-number")
}
