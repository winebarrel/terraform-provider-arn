# arn:aws:wellarchitected:ap-northeast-1:111111111111:agent-profile/profile-name
output "wellarchitected_agent_profile" {
  value = provider::arn::wellarchitected_agent_profile("profile-name")
}
