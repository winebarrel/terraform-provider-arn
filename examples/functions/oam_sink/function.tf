# arn:aws:oam:ap-northeast-1:111111111111:sink/resource-id
output "oam_sink" {
  value = provider::arn::oam_sink("resource-id")
}
