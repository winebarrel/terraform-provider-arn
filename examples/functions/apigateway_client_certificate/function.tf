# arn:aws:apigateway:ap-northeast-1::/clientcertificates/client-certificate-id
output "apigateway_client_certificate" {
  value = provider::arn::apigateway_client_certificate("client-certificate-id")
}
