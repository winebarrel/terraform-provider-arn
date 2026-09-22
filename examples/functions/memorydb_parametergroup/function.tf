# arn:aws:memorydb:ap-northeast-1:111111111111:parametergroup/parameter-group-name
output "memorydb_parametergroup" {
  value = provider::arn::memorydb_parametergroup("parameter-group-name")
}
