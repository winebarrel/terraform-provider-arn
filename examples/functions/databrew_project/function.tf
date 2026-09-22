# arn:aws:databrew:ap-northeast-1:111111111111:project/resource-id
output "databrew_project" {
  value = provider::arn::databrew_project("resource-id")
}
