# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/auth/*
output "amplifybackend_auth" {
  value = provider::arn::amplifybackend_auth("app-id")
}
