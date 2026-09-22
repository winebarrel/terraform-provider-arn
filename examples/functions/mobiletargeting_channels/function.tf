# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/channels
output "mobiletargeting_channels" {
  value = provider::arn::mobiletargeting_channels("app-id")
}
