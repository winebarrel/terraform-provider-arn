# arn:aws:securitylake:ap-northeast-1:111111111111:data-lake/default
output "securitylake_data_lake" {
  value = provider::arn::securitylake_data_lake()
}
