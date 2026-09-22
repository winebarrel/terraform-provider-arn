# arn:aws:rolesanywhere:ap-northeast-1:111111111111:subject/subject-id
output "rolesanywhere_subject" {
  value = provider::arn::rolesanywhere_subject("subject-id")
}
