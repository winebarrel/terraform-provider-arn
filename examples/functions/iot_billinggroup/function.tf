# arn:aws:iot:ap-northeast-1:111111111111:billinggroup/billing-group-name
output "iot_billinggroup" {
  value = provider::arn::iot_billinggroup("billing-group-name")
}
