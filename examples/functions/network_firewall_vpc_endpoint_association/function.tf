# arn:aws:network-firewall:ap-northeast-1:111111111111:vpc-endpoint-association/name
output "network_firewall_vpc_endpoint_association" {
  value = provider::arn::network_firewall_vpc_endpoint_association("name")
}
