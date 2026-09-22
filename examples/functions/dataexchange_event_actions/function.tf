# arn:aws:dataexchange:ap-northeast-1:111111111111:event-actions/event-action-id
output "dataexchange_event_actions" {
  value = provider::arn::dataexchange_event_actions("event-action-id")
}
