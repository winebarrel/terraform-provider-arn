# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/index/index-id/data-source/data-source-id
output "qbusiness_data_source" {
  value = provider::arn::qbusiness_data_source("application-id", "index-id", "data-source-id")
}
