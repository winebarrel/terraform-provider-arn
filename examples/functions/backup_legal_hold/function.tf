# arn:aws:backup:ap-northeast-1:111111111111:legal-hold:legal-hold-id
output "backup_legal_hold" {
  value = provider::arn::backup_legal_hold("legal-hold-id")
}
