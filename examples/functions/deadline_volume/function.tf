# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/fleet/fleet-id/volume/volume-id
output "deadline_volume" {
  value = provider::arn::deadline_volume("farm-id", "fleet-id", "volume-id")
}
