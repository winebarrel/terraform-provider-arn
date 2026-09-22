# arn:aws:omics:ap-northeast-1:111111111111:variantStore/variant-store-name
output "omics_variant_store" {
  value = provider::arn::omics_variant_store("variant-store-name")
}
