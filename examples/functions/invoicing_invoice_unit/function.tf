# arn:aws:invoicing::111111111111:invoice-unit/identifier
output "invoicing_invoice_unit" {
  value = provider::arn::invoicing_invoice_unit("identifier")
}
