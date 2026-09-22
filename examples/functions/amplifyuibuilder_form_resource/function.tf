# arn:aws:amplifyuibuilder:ap-northeast-1:111111111111:app/app-id/environment/environment-name/forms/id
output "amplifyuibuilder_form_resource" {
  value = provider::arn::amplifyuibuilder_form_resource("app-id", "environment-name", "id")
}
