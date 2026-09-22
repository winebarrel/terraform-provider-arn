# arn:aws:amplify:ap-northeast-1:111111111111:apps/app-id/branches/branch-name/jobs/job-id
output "amplify_jobs" {
  value = provider::arn::amplify_jobs("app-id", "branch-name", "job-id")
}
