# arn:aws:proton:ap-northeast-1:111111111111:service/name
output "proton_service" {
  value = provider::arn::proton_service("name")
}
