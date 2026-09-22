# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/data-accessor/data-accessor-id
output "qbusiness_data_accessor" {
  value = provider::arn::qbusiness_data_accessor("application-id", "data-accessor-id")
}
