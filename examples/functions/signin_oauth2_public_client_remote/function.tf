# arn:aws:signin:ap-northeast-1:111111111111:oauth2/public-client/remote
output "signin_oauth2_public_client_remote" {
  value = provider::arn::signin_oauth2_public_client_remote()
}
