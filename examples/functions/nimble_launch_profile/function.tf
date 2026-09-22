# arn:aws:nimble:ap-northeast-1:111111111111:launch-profile/launch-profile-id
output "nimble_launch_profile" {
  value = provider::arn::nimble_launch_profile("launch-profile-id")
}
