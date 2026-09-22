# arn:aws:omics:ap-northeast-1:111111111111:runBatch/batch-id
output "omics_run_batch" {
  value = provider::arn::omics_run_batch("batch-id")
}
