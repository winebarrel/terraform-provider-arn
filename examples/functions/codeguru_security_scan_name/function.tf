# arn:aws:codeguru-security:ap-northeast-1:111111111111:scans/scan-name
output "codeguru_security_scan_name" {
  value = provider::arn::codeguru_security_scan_name("scan-name")
}
