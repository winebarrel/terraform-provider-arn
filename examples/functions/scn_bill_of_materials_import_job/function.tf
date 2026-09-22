# arn:aws:scn:ap-northeast-1:111111111111:instance/instance-id/bill-of-materials-import-job/job-id
output "scn_bill_of_materials_import_job" {
  value = provider::arn::scn_bill_of_materials_import_job("instance-id", "job-id")
}
