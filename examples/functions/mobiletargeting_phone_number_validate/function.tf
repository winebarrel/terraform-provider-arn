# arn:aws:mobiletargeting:ap-northeast-1:111111111111:phone/number/validate
output "mobiletargeting_phone_number_validate" {
  value = provider::arn::mobiletargeting_phone_number_validate()
}
