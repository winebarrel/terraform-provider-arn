# arn:aws:ec2:ap-northeast-1:111111111111:client-vpn-endpoint/client-vpn-endpoint-id
output "ec2_client_vpn_endpoint" {
  value = provider::arn::ec2_client_vpn_endpoint("client-vpn-endpoint-id")
}
