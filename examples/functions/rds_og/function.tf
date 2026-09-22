# arn:aws:rds:ap-northeast-1:111111111111:og:option-group-name
output "rds_og" {
  value = provider::arn::rds_og("option-group-name")
}
