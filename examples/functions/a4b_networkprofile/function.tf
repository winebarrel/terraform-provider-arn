# arn:aws:a4b:ap-northeast-1:111111111111:network-profile/resource-id
output "a4b_networkprofile" {
  value = provider::arn::a4b_networkprofile("resource-id")
}
