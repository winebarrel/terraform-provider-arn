# arn:aws:amplifybackend:ap-northeast-1:111111111111:/backend/app-id/job/*
output "amplifybackend_job" {
  value = provider::arn::amplifybackend_job("app-id")
}
