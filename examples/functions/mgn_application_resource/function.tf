# arn:aws:mgn:ap-northeast-1:111111111111:application/application-id
output "mgn_application_resource" {
  value = provider::arn::mgn_application_resource("application-id")
}
