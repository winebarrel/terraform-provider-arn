# arn:aws:ec2:ap-northeast-1:111111111111:vpc-flow-log/vpc-flow-log-id
output "ec2_vpc_flow_log" {
  value = provider::arn::ec2_vpc_flow_log("vpc-flow-log-id")
}
