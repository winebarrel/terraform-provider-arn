# arn:aws:rds:ap-northeast-1:111111111111:shard-group:db-shard-group-resource-id
output "rds_shardgrp" {
  value = provider::arn::rds_shardgrp("db-shard-group-resource-id")
}
