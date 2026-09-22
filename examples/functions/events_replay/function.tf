# arn:aws:events:ap-northeast-1:111111111111:replay/replay-name
output "events_replay" {
  value = provider::arn::events_replay("replay-name")
}
