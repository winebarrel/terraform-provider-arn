# arn:aws:machinelearning:ap-northeast-1:111111111111:datasource/datasource-id
output "machinelearning_datasource" {
  value = provider::arn::machinelearning_datasource("datasource-id")
}
