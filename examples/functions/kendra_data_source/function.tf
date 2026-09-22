# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/data-source/data-source-id
output "kendra_data_source" {
  value = provider::arn::kendra_data_source("index-id", "data-source-id")
}
