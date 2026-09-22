# arn:aws:dsql:ap-northeast-1:111111111111:cluster/identifier
output "dsql_cluster" {
  value = provider::arn::dsql_cluster("identifier")
}
