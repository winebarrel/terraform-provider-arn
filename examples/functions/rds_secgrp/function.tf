# arn:aws:rds:ap-northeast-1:111111111111:secgrp:security-group-name
output "rds_secgrp" {
  value = provider::arn::rds_secgrp("security-group-name")
}
