# arn:aws:proton:ap-northeast-1:111111111111:service-template/template-name:major-version.minor-version
output "proton_service_template_version" {
  value = provider::arn::proton_service_template_version("template-name", "major-version", "minor-version")
}
