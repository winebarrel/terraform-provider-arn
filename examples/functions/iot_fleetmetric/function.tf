# arn:aws:iot:ap-northeast-1:111111111111:fleetmetric/fleet-metric-name
output "iot_fleetmetric" {
  value = provider::arn::iot_fleetmetric("fleet-metric-name")
}
