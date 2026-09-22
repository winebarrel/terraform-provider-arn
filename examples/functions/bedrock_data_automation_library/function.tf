# arn:aws:bedrock:ap-northeast-1:111111111111:data-automation-library/data-automation-library-id
output "bedrock_data_automation_library" {
  value = provider::arn::bedrock_data_automation_library("data-automation-library-id")
}
