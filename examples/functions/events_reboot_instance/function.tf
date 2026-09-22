# arn:aws:events:ap-northeast-1:111111111111:target/reboot-instance
output "events_reboot_instance" {
  value = provider::arn::events_reboot_instance()
}
