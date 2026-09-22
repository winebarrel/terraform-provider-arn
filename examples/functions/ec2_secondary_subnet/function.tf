# arn:aws:ec2:ap-northeast-1:111111111111:secondary-subnet/secondary-subnet-id
output "ec2_secondary_subnet" {
  value = provider::arn::ec2_secondary_subnet("secondary-subnet-id")
}
