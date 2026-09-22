# arn:aws:dsql:ap-northeast-1:111111111111:cluster/cluster-id/stream/stream-id
output "dsql_stream" {
  value = provider::arn::dsql_stream("cluster-id", "stream-id")
}
