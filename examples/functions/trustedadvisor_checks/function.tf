# arn:aws:trustedadvisor:ap-northeast-1:111111111111:checks/category-code/check-id
output "trustedadvisor_checks" {
  value = provider::arn::trustedadvisor_checks("category-code", "check-id")
}
