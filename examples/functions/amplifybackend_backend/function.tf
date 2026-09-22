# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/*
output "amplifybackend_backend" {
  value = provider::arn::amplifybackend_backend("app-id")
}
