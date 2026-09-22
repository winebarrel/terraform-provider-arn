# arn:aws:monitron:ap-northeast-1:111111111111:project/resource-id
output "monitron_project" {
  value = provider::arn::monitron_project("resource-id")
}
