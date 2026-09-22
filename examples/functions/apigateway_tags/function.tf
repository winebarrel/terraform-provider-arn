# arn:aws:apigateway:ap-northeast-1::/tags/url-encoded-resource-arn
output "apigateway_tags" {
  value = provider::arn::apigateway_tags("url-encoded-resource-arn")
}
