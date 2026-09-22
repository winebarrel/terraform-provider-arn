# arn:aws:cloudwatch:ap-northeast-1:111111111111:alert/alert-id
output "cloudwatch_alert" {
  value = provider::arn::cloudwatch_alert("alert-id")
}
