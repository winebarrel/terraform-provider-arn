# arn:aws:personalize:ap-northeast-1:111111111111:solution/resource-id
output "personalize_solution" {
  value = provider::arn::personalize_solution("resource-id")
}
