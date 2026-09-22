# arn:aws:ec2:ap-northeast-1:111111111111:security-group/security-group-id
output "ec2_security_group" {
  value = provider::arn::ec2_security_group("security-group-id")
}
