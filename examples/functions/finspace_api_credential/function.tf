# arn:aws:finspace-api:ap-northeast-1:111111111111:/credentials/programmatic
output "finspace_api_credential" {
  value = provider::arn::finspace_api_credential()
}
