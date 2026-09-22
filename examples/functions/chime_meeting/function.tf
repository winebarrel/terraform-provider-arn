# arn:aws:chime:ap-northeast-1:111111111111:meeting/meeting-id
output "chime_meeting" {
  value = provider::arn::chime_meeting("meeting-id")
}
