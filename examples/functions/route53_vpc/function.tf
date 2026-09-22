# arn:aws:ec2:ap-northeast-1:111111111111:vpc/vpc-id
output "route53_vpc" {
  value = provider::arn::route53_vpc("vpc-id")
}
