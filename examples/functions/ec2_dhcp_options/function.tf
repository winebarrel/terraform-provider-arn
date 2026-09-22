# arn:aws:ec2:ap-northeast-1:111111111111:dhcp-options/dhcp-options-id
output "ec2_dhcp_options" {
  value = provider::arn::ec2_dhcp_options("dhcp-options-id")
}
