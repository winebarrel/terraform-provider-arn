# arn:aws:ssm-incidents::111111111111:incident-record/response-plan/incident-record
output "ssm_incidents_incident_record" {
  value = provider::arn::ssm_incidents_incident_record("response-plan", "incident-record")
}
