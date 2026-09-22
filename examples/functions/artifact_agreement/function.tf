# arn:aws:artifact:::agreement/*
output "artifact_agreement" {
  value = provider::arn::artifact_agreement()
}
