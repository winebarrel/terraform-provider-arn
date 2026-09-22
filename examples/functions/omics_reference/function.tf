# arn:aws:omics:ap-northeast-1:111111111111:referenceStore/reference-store-id/reference/reference-id
output "omics_reference" {
  value = provider::arn::omics_reference("reference-store-id", "reference-id")
}
