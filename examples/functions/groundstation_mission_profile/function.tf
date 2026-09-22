# arn:aws:groundstation:ap-northeast-1:111111111111:mission-profile/mission-profile-id
output "groundstation_mission_profile" {
  value = provider::arn::groundstation_mission_profile("mission-profile-id")
}
