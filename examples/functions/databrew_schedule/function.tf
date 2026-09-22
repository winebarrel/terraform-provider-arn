# arn:aws:databrew:ap-northeast-1:111111111111:schedule/resource-id
output "databrew_schedule" {
  value = provider::arn::databrew_schedule("resource-id")
}
