# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/challenge/*
output "amplifybackend_token" {
  value = provider::arn::amplifybackend_token("app-id")
}
