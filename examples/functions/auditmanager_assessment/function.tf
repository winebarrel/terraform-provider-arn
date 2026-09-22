# arn:aws:auditmanager:ap-northeast-1:111111111111:assessment/assessment-id
output "auditmanager_assessment" {
  value = provider::arn::auditmanager_assessment("assessment-id")
}
