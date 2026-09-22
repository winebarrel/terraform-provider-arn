# arn:aws:inspector2:ap-northeast-1:111111111111:owner/owner-id/cis-configuration/cis-scan-configuration-id
output "inspector2_cis_scan_configuration" {
  value = provider::arn::inspector2_cis_scan_configuration("owner-id", "cis-scan-configuration-id")
}
