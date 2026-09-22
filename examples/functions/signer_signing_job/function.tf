# arn:aws:signer:ap-northeast-1:111111111111:/signing-jobs/job-id
output "signer_signing_job" {
  value = provider::arn::signer_signing_job("job-id")
}
