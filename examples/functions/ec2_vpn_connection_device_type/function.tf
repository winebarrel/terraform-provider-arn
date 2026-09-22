# arn:aws:ec2:ap-northeast-1:111111111111:vpn-connection-device-type/vpn-connection-device-type-id
output "ec2_vpn_connection_device_type" {
  value = provider::arn::ec2_vpn_connection_device_type("vpn-connection-device-type-id")
}
