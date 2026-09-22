# arn:aws:ec2:ap-northeast-1:111111111111:route-server-peer/route-server-peer-id
output "ec2_route_server_peer" {
  value = provider::arn::ec2_route_server_peer("route-server-peer-id")
}
