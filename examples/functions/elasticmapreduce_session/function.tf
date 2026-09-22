# arn:aws:elasticmapreduce:ap-northeast-1:111111111111:cluster/cluster-id/session/session-id
output "elasticmapreduce_session" {
  value = provider::arn::elasticmapreduce_session("cluster-id", "session-id")
}
