# arn:aws:omics:ap-northeast-1:111111111111:configuration/name
output "omics_configuration" {
  value = provider::arn::omics_configuration("name")
}
