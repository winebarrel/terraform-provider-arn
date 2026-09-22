# arn:aws:bedrock:ap-northeast-1:111111111111:data-automation-profile/profile-id
output "bedrock_data_automation_profile" {
  value = provider::arn::bedrock_data_automation_profile("profile-id")
}
