# arn:aws:mediaconnect:ap-northeast-1:111111111111:flow:flow-id:flow-name/vpcInterface/vpc-interface-name
output "mediaconnect_vpc_interface" {
  value = provider::arn::mediaconnect_vpc_interface("flow-id", "flow-name", "vpc-interface-name")
}
