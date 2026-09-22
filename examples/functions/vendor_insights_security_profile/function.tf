# arn:aws:vendor-insights:::security-profile:resource-id
output "vendor_insights_security_profile" {
  value = provider::arn::vendor_insights_security_profile("resource-id")
}
