# arn:aws:events:ap-northeast-1:111111111111:archive/archive-name
output "events_archive" {
  value = provider::arn::events_archive("archive-name")
}
