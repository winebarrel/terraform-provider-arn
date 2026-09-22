# arn:aws:memorydb::111111111111:multiregioncluster/cluster-name
output "memorydb_multiregioncluster" {
  value = provider::arn::memorydb_multiregioncluster("cluster-name")
}
