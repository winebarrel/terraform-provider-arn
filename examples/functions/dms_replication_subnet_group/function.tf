# arn:aws:dms:ap-northeast-1:111111111111:subgrp:*
output "dms_replication_subnet_group" {
  value = provider::arn::dms_replication_subnet_group()
}
