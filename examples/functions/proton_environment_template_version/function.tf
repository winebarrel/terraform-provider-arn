# arn:aws:proton:ap-northeast-1:111111111111:environment-template/template-name:major-version.minor-version
output "proton_environment_template_version" {
  value = provider::arn::proton_environment_template_version("template-name", "major-version", "minor-version")
}
