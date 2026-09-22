# arn:aws:lookoutmetrics:ap-northeast-1:111111111111:Alert:alert-name
output "lookoutmetrics_alert" {
  value = provider::arn::lookoutmetrics_alert("alert-name")
}
