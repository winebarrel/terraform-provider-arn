# arn:aws:omics:ap-northeast-1:111111111111:referenceStore/reference-store-id
output "omics_reference_store" {
  value = provider::arn::omics_reference_store("reference-store-id")
}
