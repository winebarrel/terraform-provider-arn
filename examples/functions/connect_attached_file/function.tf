# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/file/file-id
output "connect_attached_file" {
  value = provider::arn::connect_attached_file("instance-id", "file-id")
}
