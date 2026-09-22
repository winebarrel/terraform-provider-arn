# arn:aws:groundstation:ap-northeast-1:111111111111:satellite/satellite-id
output "groundstation_satellite" {
  value = provider::arn::groundstation_satellite("satellite-id")
}
