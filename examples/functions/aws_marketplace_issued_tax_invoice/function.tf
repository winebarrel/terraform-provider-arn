# arn:aws:aws-marketplace:ap-northeast-1:111111111111:catalog/catalog/issued-tax-invoice/resource-id
output "aws_marketplace_issued_tax_invoice" {
  value = provider::arn::aws_marketplace_issued_tax_invoice("catalog", "resource-id")
}
