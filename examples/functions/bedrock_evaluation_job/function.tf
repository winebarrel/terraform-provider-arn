# arn:aws:bedrock:ap-northeast-1:111111111111:evaluation-job/resource-id
output "bedrock_evaluation_job" {
  value = provider::arn::bedrock_evaluation_job("resource-id")
}
