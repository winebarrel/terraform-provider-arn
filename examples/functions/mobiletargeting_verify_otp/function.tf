# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/verify-otp
output "mobiletargeting_verify_otp" {
  value = provider::arn::mobiletargeting_verify_otp("app-id")
}
