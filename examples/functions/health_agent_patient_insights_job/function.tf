# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id/patient-insights-job/job-id
output "health_agent_patient_insights_job" {
  value = provider::arn::health_agent_patient_insights_job("domain-id", "job-id")
}
