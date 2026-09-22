# arn:aws:secretsmanager:ap-northeast-1:111111111111:secret:secret-id
output "secretsmanager_secret" {
  value = provider::arn::secretsmanager_secret("secret-id")
}
