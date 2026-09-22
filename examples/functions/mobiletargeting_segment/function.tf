# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/segments/segment-id
output "mobiletargeting_segment" {
  value = provider::arn::mobiletargeting_segment("app-id", "segment-id")
}
