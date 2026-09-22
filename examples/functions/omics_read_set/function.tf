# arn:aws:omics:ap-northeast-1:111111111111:sequenceStore/sequence-store-id/readSet/read-set-id
output "omics_read_set" {
  value = provider::arn::omics_read_set("sequence-store-id", "read-set-id")
}
