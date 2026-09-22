# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/api/*
output "amplifybackend_api" {
  value = provider::arn::amplifybackend_api("app-id")
}
