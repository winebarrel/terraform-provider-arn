# arn:aws:redshift:ap-northeast-1:111111111111:securitygroup:security-group-name/ec2securitygroup/owner/ec2-security-group-id
output "redshift_securitygroup" {
  value = provider::arn::redshift_securitygroup("security-group-name", "owner", "ec2-security-group-id")
}
