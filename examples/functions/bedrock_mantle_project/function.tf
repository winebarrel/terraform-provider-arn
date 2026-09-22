# arn:aws:bedrock-mantle:ap-northeast-1:111111111111:project/resource-id
output "bedrock_mantle_project" {
  value = provider::arn::bedrock_mantle_project("resource-id")
}
