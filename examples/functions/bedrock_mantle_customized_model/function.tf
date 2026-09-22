# arn:aws:bedrock-mantle:ap-northeast-1:111111111111:customized-model/resource-id
output "bedrock_mantle_customized_model" {
  value = provider::arn::bedrock_mantle_customized_model("resource-id")
}
