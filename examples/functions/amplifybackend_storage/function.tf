# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/storage/*
output "amplifybackend_storage" {
  value = provider::arn::amplifybackend_storage("app-id")
}
