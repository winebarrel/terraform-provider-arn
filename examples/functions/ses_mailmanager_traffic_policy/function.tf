# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-traffic-policy/traffic-policy-id
output "ses_mailmanager_traffic_policy" {
  value = provider::arn::ses_mailmanager_traffic_policy("traffic-policy-id")
}
