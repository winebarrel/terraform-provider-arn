# arn:aws:auditmanager:ap-northeast-1:111111111111:assessmentFramework/assessment-framework-id
output "auditmanager_assessment_framework" {
  value = provider::arn::auditmanager_assessment_framework("assessment-framework-id")
}
