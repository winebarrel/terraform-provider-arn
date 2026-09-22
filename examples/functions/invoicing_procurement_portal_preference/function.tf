# arn:aws:invoicing::111111111111:procurement-portal-preference/identifier
output "invoicing_procurement_portal_preference" {
  value = provider::arn::invoicing_procurement_portal_preference("identifier")
}
