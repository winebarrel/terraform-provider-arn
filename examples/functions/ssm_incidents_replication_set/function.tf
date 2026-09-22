# arn:aws:ssm-incidents::111111111111:replication-set/replication-set
output "ssm_incidents_replication_set" {
  value = provider::arn::ssm_incidents_replication_set("replication-set")
}
