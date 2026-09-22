# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/fleet/fleet-id/worker/worker-id
output "deadline_worker" {
  value = provider::arn::deadline_worker("farm-id", "fleet-id", "worker-id")
}
