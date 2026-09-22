# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/endpoints/endpoint-id
output "mobiletargeting_endpoint" {
  value = provider::arn::mobiletargeting_endpoint("app-id", "endpoint-id")
}
