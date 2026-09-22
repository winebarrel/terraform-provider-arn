# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id
output "connect_instance" {
  value = provider::arn::connect_instance("instance-id")
}
