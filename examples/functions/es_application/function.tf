# arn:aws:opensearch:ap-northeast-1:111111111111:application/app-id
output "es_application" {
  value = provider::arn::es_application("app-id")
}
