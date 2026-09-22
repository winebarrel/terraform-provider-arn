# arn:aws:groundstation:ap-northeast-1:111111111111:groundstation:ground-station-id
output "groundstation_ground_station_resource" {
  value = provider::arn::groundstation_ground_station_resource("ground-station-id")
}
