# arn:aws:cloudwatch:ap-northeast-1:111111111111:space/space-id
output "cloudwatch_space" {
  value = provider::arn::cloudwatch_space("space-id")
}
