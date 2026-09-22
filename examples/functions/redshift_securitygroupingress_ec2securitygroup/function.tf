# arn:aws:redshift:ap-northeast-1:111111111111:securitygroupingress:security-group-name/ec2securitygroup/owner/ece2-securitygroup-id
output "redshift_securitygroupingress_ec2securitygroup" {
  value = provider::arn::redshift_securitygroupingress_ec2securitygroup("security-group-name", "owner", "ece2-securitygroup-id")
}
