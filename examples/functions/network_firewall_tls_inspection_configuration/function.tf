# arn:aws:network-firewall:ap-northeast-1:111111111111:tls-configuration/name
output "network_firewall_tls_inspection_configuration" {
  value = provider::arn::network_firewall_tls_inspection_configuration("name")
}
