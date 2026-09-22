# arn:aws:mediaconvert:ap-northeast-1:111111111111:jobTemplates/job-template-name
output "mediaconvert_job_template" {
  value = provider::arn::mediaconvert_job_template("job-template-name")
}
