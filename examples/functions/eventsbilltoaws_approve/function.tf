# arn:aws:eventsbilltoaws:ap-northeast-1:111111111111:relative-id
output "eventsbilltoaws_approve" {
  value = provider::arn::eventsbilltoaws_approve("relative-id")
}
