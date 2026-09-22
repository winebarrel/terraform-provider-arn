# arn:aws:a4b:ap-northeast-1:111111111111:schedule/resource-id
output "a4b_schedule" {
  value = provider::arn::a4b_schedule("resource-id")
}
