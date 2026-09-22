# arn:aws:omics:ap-northeast-1:111111111111:annotationStore/annotation-store-name/version/annotation-store-version-name
output "omics_annotation_store_version" {
  value = provider::arn::omics_annotation_store_version("annotation-store-name", "annotation-store-version-name")
}
