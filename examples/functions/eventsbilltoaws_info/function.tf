# arn:aws:eventsbilltoaws:ap-northeast-1:111111111111:relative-id
output "eventsbilltoaws_info" {
  value = provider::arn::eventsbilltoaws_info("relative-id")
}
