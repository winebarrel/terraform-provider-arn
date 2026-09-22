# arn:aws:translate:ap-northeast-1:111111111111:parallel-data/resource-name
output "translate_parallel_data" {
  value = provider::arn::translate_parallel_data("resource-name")
}
