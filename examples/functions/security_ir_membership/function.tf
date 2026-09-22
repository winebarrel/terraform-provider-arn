# arn:aws:security-ir:ap-northeast-1:111111111111:membership/membership-id
output "security_ir_membership" {
  value = provider::arn::security_ir_membership("membership-id")
}
