# arn:aws:macie2:ap-northeast-1:111111111111:allow-list/resource-id
output "macie2_allow_list" {
  value = provider::arn::macie2_allow_list("resource-id")
}
