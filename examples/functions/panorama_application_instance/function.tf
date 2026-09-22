# arn:aws:panorama:ap-northeast-1:111111111111:applicationInstance/application-instance-id
output "panorama_application_instance" {
  value = provider::arn::panorama_application_instance("application-instance-id")
}
