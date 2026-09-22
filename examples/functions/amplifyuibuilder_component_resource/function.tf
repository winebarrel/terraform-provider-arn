# arn:aws:amplifyuibuilder:ap-northeast-1:111111111111:app/app-id/environment/environment-name/components/id
output "amplifyuibuilder_component_resource" {
  value = provider::arn::amplifyuibuilder_component_resource("app-id", "environment-name", "id")
}
