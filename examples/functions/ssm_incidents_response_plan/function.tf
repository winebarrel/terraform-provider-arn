# arn:aws:ssm-incidents::111111111111:response-plan/response-plan
output "ssm_incidents_response_plan" {
  value = provider::arn::ssm_incidents_response_plan("response-plan")
}
