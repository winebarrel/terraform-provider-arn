# arn:aws:gamelift:ap-northeast-1:111111111111:containerfleet/fleet-id
output "gamelift_container_fleet" {
  value = provider::arn::gamelift_container_fleet("fleet-id")
}
