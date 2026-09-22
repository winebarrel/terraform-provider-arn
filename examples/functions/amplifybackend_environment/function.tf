# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/environments/*
output "amplifybackend_environment" {
  value = provider::arn::amplifybackend_environment("app-id")
}
