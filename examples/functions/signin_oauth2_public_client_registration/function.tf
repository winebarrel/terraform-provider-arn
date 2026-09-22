# arn:aws:signin:ap-northeast-1::external-client/dcr/*
output "signin_oauth2_public_client_registration" {
  value = provider::arn::signin_oauth2_public_client_registration()
}
