# arn:aws:scn:ap-northeast-1:111111111111:instance/instance-id
output "scn_instance" {
  value = provider::arn::scn_instance("instance-id")
}
