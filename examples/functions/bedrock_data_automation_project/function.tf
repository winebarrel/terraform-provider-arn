# arn:aws:bedrock:ap-northeast-1:111111111111:data-automation-project/project-id
output "bedrock_data_automation_project" {
  value = provider::arn::bedrock_data_automation_project("project-id")
}
