# arn:aws:network-firewall:ap-northeast-1:111111111111:container-association/name
output "network_firewall_container_association" {
  value = provider::arn::network_firewall_container_association("name")
}
