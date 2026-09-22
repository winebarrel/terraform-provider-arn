# arn:aws:bedrock:ap-northeast-1:111111111111:project/resource-id
output "bedrock_project" {
  value = provider::arn::bedrock_project("resource-id")
}
