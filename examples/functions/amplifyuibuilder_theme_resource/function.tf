# arn:aws:amplifyuibuilder:ap-northeast-1:111111111111:app/app-id/environment/environment-name/themes/id
output "amplifyuibuilder_theme_resource" {
  value = provider::arn::amplifyuibuilder_theme_resource("app-id", "environment-name", "id")
}
