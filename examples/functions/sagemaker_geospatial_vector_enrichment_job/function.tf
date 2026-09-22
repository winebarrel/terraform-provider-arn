# arn:aws:sagemaker-geospatial:ap-northeast-1:111111111111:vector-enrichment-job/job-id
output "sagemaker_geospatial_vector_enrichment_job" {
  value = provider::arn::sagemaker_geospatial_vector_enrichment_job("job-id")
}
