# arn:aws:ec2:ap-northeast-1:111111111111:route-server/route-server-id
output "ec2_route_server" {
  value = provider::arn::ec2_route_server("route-server-id")
}
