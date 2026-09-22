# arn:aws:workmailmessageflow:ap-northeast-1:111111111111:message/organization-id/context/message-id
output "workmailmessageflow_raw_message" {
  value = provider::arn::workmailmessageflow_raw_message("organization-id", "context", "message-id")
}
