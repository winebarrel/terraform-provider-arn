# arn:aws:docdb-elastic:ap-northeast-1:111111111111:cluster/resource-id
output "docdb_elastic_cluster" {
  value = provider::arn::docdb_elastic_cluster("resource-id")
}
