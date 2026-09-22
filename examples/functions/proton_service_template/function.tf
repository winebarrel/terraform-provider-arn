# arn:aws:proton:ap-northeast-1:111111111111:service-template/name
output "proton_service_template" {
  value = provider::arn::proton_service_template("name")
}
