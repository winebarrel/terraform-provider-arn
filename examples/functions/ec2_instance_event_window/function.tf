# arn:aws:ec2:ap-northeast-1:111111111111:instance-event-window/instance-event-window-id
output "ec2_instance_event_window" {
  value = provider::arn::ec2_instance_event_window("instance-event-window-id")
}
