# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:audience-generation-job/resource-id
output "cleanrooms_ml_audiencegenerationjob" {
  value = provider::arn::cleanrooms_ml_audiencegenerationjob("resource-id")
}
