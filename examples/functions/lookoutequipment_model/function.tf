# arn:aws:lookoutequipment:ap-northeast-1:111111111111:model/model-name/model-id
output "lookoutequipment_model" {
  value = provider::arn::lookoutequipment_model("model-name", "model-id")
}
