# arn:aws:artifact::111111111111:customer-agreement/*
output "artifact_customer_agreement" {
  value = provider::arn::artifact_customer_agreement()
}
