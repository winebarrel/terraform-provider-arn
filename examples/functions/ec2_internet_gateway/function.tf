# arn:aws:ec2:ap-northeast-1:111111111111:internet-gateway/internet-gateway-id
output "ec2_internet_gateway" {
  value = provider::arn::ec2_internet_gateway("internet-gateway-id")
}
