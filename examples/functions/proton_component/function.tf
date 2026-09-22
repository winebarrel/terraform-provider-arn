# arn:aws:proton:ap-northeast-1:111111111111:component/id
output "proton_component" {
  value = provider::arn::proton_component("id")
}
