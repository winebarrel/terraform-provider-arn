# arn:aws:security-ir:ap-northeast-1:111111111111:case/case-id
output "security_ir_case" {
  value = provider::arn::security_ir_case("case-id")
}
