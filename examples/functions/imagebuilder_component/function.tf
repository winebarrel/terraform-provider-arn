# arn:aws:imagebuilder:ap-northeast-1:111111111111:component/component-name/component-version/component-build-version
output "imagebuilder_component" {
  value = provider::arn::imagebuilder_component("component-name", "component-version", "component-build-version")
}
