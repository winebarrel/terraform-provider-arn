# arn:aws:imagebuilder:ap-northeast-1:111111111111:component/component-name/component-version/*
output "imagebuilder_all_component_build_versions" {
  value = provider::arn::imagebuilder_all_component_build_versions("component-name", "component-version")
}
