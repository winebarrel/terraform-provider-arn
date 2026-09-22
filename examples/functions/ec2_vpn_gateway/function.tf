# arn:aws:ec2:ap-northeast-1:111111111111:vpn-gateway/vpn-gateway-id
output "ec2_vpn_gateway" {
  value = provider::arn::ec2_vpn_gateway("vpn-gateway-id")
}
