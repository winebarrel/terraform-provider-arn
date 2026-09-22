# arn:aws:auditmanager:ap-northeast-1:111111111111:assessment/assessment-id/controlSet/control-set-id
output "auditmanager_assessment_control_set" {
  value = provider::arn::auditmanager_assessment_control_set("assessment-id", "control-set-id")
}
