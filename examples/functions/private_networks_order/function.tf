# arn:aws:private-networks:ap-northeast-1:111111111111:order/network-name/order-id
output "private_networks_order" {
  value = provider::arn::private_networks_order("network-name", "order-id")
}
