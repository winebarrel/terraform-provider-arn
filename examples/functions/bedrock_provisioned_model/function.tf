# arn:aws:bedrock:ap-northeast-1:111111111111:provisioned-model/resource-id
output "bedrock_provisioned_model" {
  value = provider::arn::bedrock_provisioned_model("resource-id")
}
