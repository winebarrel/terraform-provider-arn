# arn:aws:signin:ap-northeast-1:111111111111:oauth2/public-client/localhost
output "signin_oauth2_public_client_localhost" {
  value = provider::arn::signin_oauth2_public_client_localhost()
}
