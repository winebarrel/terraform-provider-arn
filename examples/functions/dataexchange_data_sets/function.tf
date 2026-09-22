# arn:aws:dataexchange:ap-northeast-1:111111111111:data-sets/data-set-id
output "dataexchange_data_sets" {
  value = provider::arn::dataexchange_data_sets("data-set-id")
}
