# arn:aws:macie2:ap-northeast-1:111111111111:member/resource-id
output "macie2_member" {
  value = provider::arn::macie2_member("resource-id")
}
