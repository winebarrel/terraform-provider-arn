# arn:aws:codeguru-reviewer:ap-northeast-1:111111111111:association:resource-id:codereview:code-review-id
output "codeguru_reviewer_codereview" {
  value = provider::arn::codeguru_reviewer_codereview("resource-id", "code-review-id")
}
