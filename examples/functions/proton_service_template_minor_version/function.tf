# arn:aws:proton:ap-northeast-1:111111111111:service-template/template-name:major-version-id.minor-version-id
output "proton_service_template_minor_version" {
  value = provider::arn::proton_service_template_minor_version("template-name", "major-version-id", "minor-version-id")
}
