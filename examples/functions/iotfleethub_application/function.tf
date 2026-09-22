# arn:aws:iotfleethub:ap-northeast-1:111111111111:application/application-id
output "iotfleethub_application" {
  value = provider::arn::iotfleethub_application("application-id")
}
