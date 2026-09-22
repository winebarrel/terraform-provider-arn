# arn:aws:ec2:ap-northeast-1:111111111111:vpn-concentrator/vpn-concentrator-id
output "ec2_vpn_concentrator" {
  value = provider::arn::ec2_vpn_concentrator("vpn-concentrator-id")
}
