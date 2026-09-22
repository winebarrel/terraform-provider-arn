# arn:aws:apigateway:ap-northeast-1::/vpclinks/vpc-link-id
output "apigateway_vpc_link" {
  value = provider::arn::apigateway_vpc_link("vpc-link-id")
}
