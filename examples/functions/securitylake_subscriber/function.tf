# arn:aws:securitylake:ap-northeast-1:111111111111:subscriber/subscriber-id
output "securitylake_subscriber" {
  value = provider::arn::securitylake_subscriber("subscriber-id")
}
