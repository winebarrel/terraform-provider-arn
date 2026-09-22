# arn:aws:redshift:ap-northeast-1:111111111111:securitygroupingress:security-group-name/cidrip/ip-range
output "redshift_securitygroupingress_cidr" {
  value = provider::arn::redshift_securitygroupingress_cidr("security-group-name", "ip-range")
}
