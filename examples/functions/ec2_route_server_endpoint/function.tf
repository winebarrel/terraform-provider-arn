# arn:aws:ec2:ap-northeast-1:111111111111:route-server-endpoint/route-server-endpoint-id
output "ec2_route_server_endpoint" {
  value = provider::arn::ec2_route_server_endpoint("route-server-endpoint-id")
}
