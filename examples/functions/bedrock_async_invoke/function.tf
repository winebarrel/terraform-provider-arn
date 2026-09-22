# arn:aws:bedrock:ap-northeast-1:111111111111:async-invoke/resource-id
output "bedrock_async_invoke" {
  value = provider::arn::bedrock_async_invoke("resource-id")
}
