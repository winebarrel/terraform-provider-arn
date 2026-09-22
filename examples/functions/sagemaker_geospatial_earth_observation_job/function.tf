# arn:aws:sagemaker-geospatial:ap-northeast-1:111111111111:earth-observation-job/job-id
output "sagemaker_geospatial_earth_observation_job" {
  value = provider::arn::sagemaker_geospatial_earth_observation_job("job-id")
}
