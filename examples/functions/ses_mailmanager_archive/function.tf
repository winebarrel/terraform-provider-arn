# arn:aws:ses:ap-northeast-1:111111111111:mailmanager-archive/archive-id
output "ses_mailmanager_archive" {
  value = provider::arn::ses_mailmanager_archive("archive-id")
}
