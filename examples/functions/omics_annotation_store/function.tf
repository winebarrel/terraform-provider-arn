# arn:aws:omics:ap-northeast-1:111111111111:annotationStore/annotation-store-name
output "omics_annotation_store" {
  value = provider::arn::omics_annotation_store("annotation-store-name")
}
