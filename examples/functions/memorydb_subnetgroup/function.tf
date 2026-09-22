# arn:aws:memorydb:ap-northeast-1:111111111111:subnetgroup/subnet-group-name
output "memorydb_subnetgroup" {
  value = provider::arn::memorydb_subnetgroup("subnet-group-name")
}
