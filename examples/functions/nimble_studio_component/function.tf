# arn:aws:nimble:ap-northeast-1:111111111111:studio-component/studio-component-id
output "nimble_studio_component" {
  value = provider::arn::nimble_studio_component("studio-component-id")
}
