# arn:aws:ec2:ap-northeast-1:111111111111:route-table/route-table-id
output "ec2_route_table" {
  value = provider::arn::ec2_route_table("route-table-id")
}
