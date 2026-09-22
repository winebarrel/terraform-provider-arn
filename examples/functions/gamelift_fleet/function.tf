# arn:aws:gamelift:ap-northeast-1:111111111111:fleet/fleet-id
output "gamelift_fleet" {
  value = provider::arn::gamelift_fleet("fleet-id")
}
