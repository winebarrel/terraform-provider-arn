# arn:aws:amplifyuibuilder:ap-northeast-1:111111111111:app/app-id/environment/environment-name/codegen-jobs/id
output "amplifyuibuilder_codegen_job_resource" {
  value = provider::arn::amplifyuibuilder_codegen_job_resource("app-id", "environment-name", "id")
}
