# arn:aws:aws-marketplace:ap-northeast-1:111111111111:catalog/catalog/invoice-submission-task/resource-id
output "aws_marketplace_invoice_submission_task" {
  value = provider::arn::aws_marketplace_invoice_submission_task("catalog", "resource-id")
}
