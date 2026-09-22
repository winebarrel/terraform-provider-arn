# arn:aws:iotwireless:ap-northeast-1:111111111111:FuotaTask/fuota-task-id
output "iotwireless_fuota_task" {
  value = provider::arn::iotwireless_fuota_task("fuota-task-id")
}
