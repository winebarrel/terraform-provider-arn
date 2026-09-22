# arn:aws:quicksight:ap-northeast-1:111111111111:dlpsetting/resource-id
output "quicksight_dlp_setting" {
  value = provider::arn::quicksight_dlp_setting("resource-id")
}
