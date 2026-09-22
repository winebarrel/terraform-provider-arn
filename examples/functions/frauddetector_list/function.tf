# arn:aws:frauddetector:ap-northeast-1:111111111111:list/resource-path
output "frauddetector_list" {
  value = provider::arn::frauddetector_list("resource-path")
}
