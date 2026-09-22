# arn:aws:network-firewall:ap-northeast-1:111111111111:proxy-configuration/name
output "network_firewall_proxy_configuration" {
  value = provider::arn::network_firewall_proxy_configuration("name")
}
