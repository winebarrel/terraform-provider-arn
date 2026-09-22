# arn:aws:appstream:ap-northeast-1:111111111111:fleet/fleet-name
output "appstream_fleet" {
  value = provider::arn::appstream_fleet("fleet-name")
}
