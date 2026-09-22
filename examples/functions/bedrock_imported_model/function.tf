# arn:aws:bedrock:ap-northeast-1:111111111111:imported-model/resource-id
output "bedrock_imported_model" {
  value = provider::arn::bedrock_imported_model("resource-id")
}
