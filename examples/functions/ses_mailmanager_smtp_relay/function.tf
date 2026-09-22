# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-smtp-relay/relay-id
output "ses_mailmanager_smtp_relay" {
  value = provider::arn::ses_mailmanager_smtp_relay("relay-id")
}
