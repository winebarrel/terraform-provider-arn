# arn:aws:dataexchange:ap-northeast-1::data-sets/data-set-id
output "dataexchange_entitled_data_sets" {
  value = provider::arn::dataexchange_entitled_data_sets("data-set-id")
}
