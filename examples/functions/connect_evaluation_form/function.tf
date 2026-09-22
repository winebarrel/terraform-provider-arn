# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/evaluation-form/form-id
output "connect_evaluation_form" {
  value = provider::arn::connect_evaluation_form("instance-id", "form-id")
}
