# arn:aws:bedrock:ap-northeast-1:111111111111:model-import-job/resource-id
output "bedrock_model_import_job" {
  value = provider::arn::bedrock_model_import_job("resource-id")
}
