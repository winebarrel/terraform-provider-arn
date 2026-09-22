# arn:aws:lambda:ap-northeast-1:111111111111:code-signing-config:code-signing-config-id
output "lambda_code_signing_config" {
  value = provider::arn::lambda_code_signing_config("code-signing-config-id")
}
