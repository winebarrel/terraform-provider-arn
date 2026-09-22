# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/config/*
output "amplifybackend_config" {
  value = provider::arn::amplifybackend_config("app-id")
}
