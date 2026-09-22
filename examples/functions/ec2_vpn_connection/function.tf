# arn:aws:ec2:ap-northeast-1:111111111111:vpn-connection/vpn-connection-id
output "ec2_vpn_connection" {
  value = provider::arn::ec2_vpn_connection("vpn-connection-id")
}
