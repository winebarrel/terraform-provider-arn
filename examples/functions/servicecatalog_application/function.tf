# arn:aws:servicecatalog:ap-northeast-1:111111111111:/applications/application-id
output "servicecatalog_application" {
  value = provider::arn::servicecatalog_application("application-id")
}
