# arn:aws:codeguru-reviewer:ap-northeast-1:111111111111:association:resource-id
output "codeguru_reviewer_association" {
  value = provider::arn::codeguru_reviewer_association("resource-id")
}
