# arn:aws:iotwireless:ap-northeast-1:111111111111:WirelessGatewayTaskDefinition/wireless-gateway-task-definition-id
output "iotwireless_wireless_gateway_task_definition" {
  value = provider::arn::iotwireless_wireless_gateway_task_definition("wireless-gateway-task-definition-id")
}
