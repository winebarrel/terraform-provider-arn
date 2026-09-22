# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id
output "mobiletargeting_app" {
  value = provider::arn::mobiletargeting_app("app-id")
}
