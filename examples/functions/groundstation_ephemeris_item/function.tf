# arn:aws:groundstation:ap-northeast-1:111111111111:ephemeris/ephemeris-id
output "groundstation_ephemeris_item" {
  value = provider::arn::groundstation_ephemeris_item("ephemeris-id")
}
