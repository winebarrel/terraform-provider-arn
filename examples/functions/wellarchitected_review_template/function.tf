# arn:aws:wellarchitected:ap-northeast-1:111111111111:review-template/resource-id
output "wellarchitected_review_template" {
  value = provider::arn::wellarchitected_review_template("resource-id")
}
