# arn:aws:ec2:ap-northeast-1:111111111111:vpc-encryption-control/vpc-encryption-control-id
output "ec2_vpc_encryption_control" {
  value = provider::arn::ec2_vpc_encryption_control("vpc-encryption-control-id")
}
