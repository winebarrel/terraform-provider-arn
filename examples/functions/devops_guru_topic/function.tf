# arn:aws:sns:ap-northeast-1:111111111111:topic-name
output "devops_guru_topic" {
  value = provider::arn::devops_guru_topic("topic-name")
}
