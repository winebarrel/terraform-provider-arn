# arn:aws:ec2:ap-northeast-1:111111111111:egress-only-internet-gateway/egress-only-internet-gateway-id
output "ec2_egress_only_internet_gateway" {
  value = provider::arn::ec2_egress_only_internet_gateway("egress-only-internet-gateway-id")
}
