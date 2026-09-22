# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/*
output "amplifybackend_created_backend" {
  value = provider::arn::amplifybackend_created_backend()
}
