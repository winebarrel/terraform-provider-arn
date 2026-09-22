# arn:aws:inspector2:ap-northeast-1:111111111111:owner/owner-id/codesecurity-configuration/code-security-scan-configuration-id
output "inspector2_code_security_scan_configuration" {
  value = provider::arn::inspector2_code_security_scan_configuration("owner-id", "code-security-scan-configuration-id")
}
