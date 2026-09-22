# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/routing-profile/routing-profile-id
output "connect_routing_profile" {
  value = provider::arn::connect_routing_profile("instance-id", "routing-profile-id")
}
