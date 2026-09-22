# arn:aws:rds:ap-northeast-1:111111111111:subgrp:subnet-group-name
output "rds_subgrp" {
  value = provider::arn::rds_subgrp("subnet-group-name")
}
