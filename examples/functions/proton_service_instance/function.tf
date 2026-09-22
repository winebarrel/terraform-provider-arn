# arn:aws:proton:ap-northeast-1:111111111111:service/service-name/service-instance/name
output "proton_service_instance" {
  value = provider::arn::proton_service_instance("service-name", "name")
}
