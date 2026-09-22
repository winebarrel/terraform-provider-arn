# arn:aws:personalize:ap-northeast-1:111111111111:recommender/resource-id
output "personalize_recommender" {
  value = provider::arn::personalize_recommender("resource-id")
}
