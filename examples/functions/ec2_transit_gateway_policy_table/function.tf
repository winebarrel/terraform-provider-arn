# arn:aws:ec2:ap-northeast-1:111111111111:transit-gateway-policy-table/transit-gateway-policy-table-id
output "ec2_transit_gateway_policy_table" {
  value = provider::arn::ec2_transit_gateway_policy_table("transit-gateway-policy-table-id")
}
