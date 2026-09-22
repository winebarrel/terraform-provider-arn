# arn:aws:opensearch:ap-northeast-1:111111111111:application/app-id
output "opensearch_application" {
  value = provider::arn::opensearch_application("app-id")
}
