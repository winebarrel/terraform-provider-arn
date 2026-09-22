# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/vocabulary/vocabulary-id
output "connect_vocabulary" {
  value = provider::arn::connect_vocabulary("instance-id", "vocabulary-id")
}
