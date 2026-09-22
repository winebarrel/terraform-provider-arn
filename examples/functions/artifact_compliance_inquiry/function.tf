# arn:aws:artifact:ap-northeast-1:111111111111:compliance-inquiry/*
output "artifact_compliance_inquiry" {
  value = provider::arn::artifact_compliance_inquiry()
}
