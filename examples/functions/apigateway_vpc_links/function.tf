# arn:aws:apigateway:ap-northeast-1::/vpclinks
output "apigateway_vpc_links" {
  value = provider::arn::apigateway_vpc_links()
}
