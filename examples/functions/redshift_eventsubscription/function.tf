# arn:aws:redshift:ap-northeast-1:111111111111:eventsubscription:event-subscription-name
output "redshift_eventsubscription" {
  value = provider::arn::redshift_eventsubscription("event-subscription-name")
}
