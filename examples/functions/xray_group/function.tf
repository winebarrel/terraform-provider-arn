# arn:aws:xray:ap-northeast-1:111111111111:group/group-name/id
output "xray_group" {
  value = provider::arn::xray_group("group-name", "id")
}
