# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/template/template-id
output "cases_template" {
  value = provider::arn::cases_template("domain-id", "template-id")
}
