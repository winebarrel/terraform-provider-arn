# arn:aws:medialive:ap-northeast-1:111111111111:signal-map:signal-map-id
output "medialive_signal_map" {
  value = provider::arn::medialive_signal_map("signal-map-id")
}
