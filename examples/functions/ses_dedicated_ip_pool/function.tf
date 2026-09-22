# arn:aws:ses:ap-northeast-1:111111111111:dedicated-ip-pool/dedicated-ip-pool
output "ses_dedicated_ip_pool" {
  value = provider::arn::ses_dedicated_ip_pool("dedicated-ip-pool")
}
