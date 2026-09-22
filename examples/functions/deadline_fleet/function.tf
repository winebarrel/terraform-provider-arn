# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/fleet/fleet-id
output "deadline_fleet" {
  value = provider::arn::deadline_fleet("farm-id", "fleet-id")
}
