# arn:aws:iq:ap-northeast-1::company/company-id
output "iq_company" {
  value = provider::arn::iq_company("company-id")
}
