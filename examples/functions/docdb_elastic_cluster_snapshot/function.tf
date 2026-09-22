# arn:aws:docdb-elastic:ap-northeast-1:111111111111:cluster-snapshot/resource-id
output "docdb_elastic_cluster_snapshot" {
  value = provider::arn::docdb_elastic_cluster_snapshot("resource-id")
}
