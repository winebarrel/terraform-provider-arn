# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-ingress-point/ingress-point-id
output "ses_mailmanager_ingress_point" {
  value = provider::arn::ses_mailmanager_ingress_point("ingress-point-id")
}
