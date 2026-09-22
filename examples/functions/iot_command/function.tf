# arn:aws:iot:ap-northeast-1:111111111111:command/command-id
output "iot_command" {
  value = provider::arn::iot_command("command-id")
}
