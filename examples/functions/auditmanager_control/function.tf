# arn:aws:auditmanager:ap-northeast-1:111111111111:control/control-id
output "auditmanager_control" {
  value = provider::arn::auditmanager_control("control-id")
}
