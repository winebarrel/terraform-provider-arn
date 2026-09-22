# arn:aws:a4b:ap-northeast-1:111111111111:skill-group/resource-id
output "a4b_skillgroup" {
  value = provider::arn::a4b_skillgroup("resource-id")
}
