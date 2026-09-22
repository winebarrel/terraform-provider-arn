# arn:aws:quicksight:ap-northeast-1:111111111111:oauthClientApplication/resource-id
output "quicksight_oauth_client_application" {
  value = provider::arn::quicksight_oauth_client_application("resource-id")
}
