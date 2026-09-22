# arn:aws:groundstation:ap-northeast-1:111111111111:config/config-type/config-id
output "groundstation_config" {
  value = provider::arn::groundstation_config("config-type", "config-id")
}
