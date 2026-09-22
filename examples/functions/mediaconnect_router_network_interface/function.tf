# arn:aws:mediaconnect:ap-northeast-1:111111111111:routerNetworkInterface:router-network-interface-id
output "mediaconnect_router_network_interface" {
  value = provider::arn::mediaconnect_router_network_interface("router-network-interface-id")
}
