# arn:aws:memorydb::111111111111:multiregionparametergroup/multi-region-parameter-group-name
output "memorydb_multiregionparametergroup" {
  value = provider::arn::memorydb_multiregionparametergroup("multi-region-parameter-group-name")
}
