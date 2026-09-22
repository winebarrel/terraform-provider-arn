# arn:aws:ec2:ap-northeast-1:111111111111:subnet/subnet-id
output "ec2_subnet" {
  value = provider::arn::ec2_subnet("subnet-id")
}
