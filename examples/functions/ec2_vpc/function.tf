# arn:aws:ec2:ap-northeast-1:111111111111:vpc/vpc-id
output "ec2_vpc" {
  value = provider::arn::ec2_vpc("vpc-id")
}
