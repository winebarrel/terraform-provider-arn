# arn:aws:resiliencehub:ap-northeast-1:111111111111:system/system-id
output "resiliencehub_system" {
  value = provider::arn::resiliencehub_system("system-id")
}
