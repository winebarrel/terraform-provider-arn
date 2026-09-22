# arn:aws:omics:ap-northeast-1:111111111111:sequenceStore/sequence-store-id
output "omics_sequence_store" {
  value = provider::arn::omics_sequence_store("sequence-store-id")
}
