# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/contact-evaluation/evaluation-id
output "connect_contact_evaluation" {
  value = provider::arn::connect_contact_evaluation("instance-id", "evaluation-id")
}
