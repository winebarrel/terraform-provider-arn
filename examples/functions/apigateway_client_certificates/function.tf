# arn:aws:apigateway:ap-northeast-1::/clientcertificates
output "apigateway_client_certificates" {
  value = provider::arn::apigateway_client_certificates()
}
