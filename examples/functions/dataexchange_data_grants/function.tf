# arn:aws:dataexchange:ap-northeast-1:111111111111:data-grants/data-grant-id
output "dataexchange_data_grants" {
  value = provider::arn::dataexchange_data_grants("data-grant-id")
}
