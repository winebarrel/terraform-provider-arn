# arn:aws:apigateway:ap-northeast-1::/apis/api-id/stages/stage-name/accesslogsettings
output "apigateway_access_log_settings" {
  value = provider::arn::apigateway_access_log_settings("api-id", "stage-name")
}
