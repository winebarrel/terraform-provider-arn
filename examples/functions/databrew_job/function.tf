# arn:aws:databrew:ap-northeast-1:111111111111:job/resource-id
output "databrew_job" {
  value = provider::arn::databrew_job("resource-id")
}
