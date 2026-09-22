# arn:aws:lookoutequipment:ap-northeast-1:111111111111:label-group/label-group-name/label-group-id
output "lookoutequipment_label_group" {
  value = provider::arn::lookoutequipment_label_group("label-group-name", "label-group-id")
}
