# arn:aws:ec2:ap-northeast-1:111111111111:natgateway/nat-gateway-id
output "ec2_natgateway" {
  value = provider::arn::ec2_natgateway("nat-gateway-id")
}
