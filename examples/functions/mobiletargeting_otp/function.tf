# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/otp
output "mobiletargeting_otp" {
  value = provider::arn::mobiletargeting_otp("app-id")
}
