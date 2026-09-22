# arn:aws:fms:ap-northeast-1:111111111111:protocols-list/id
output "fms_protocols_list" {
  value = provider::arn::fms_protocols_list("id")
}
