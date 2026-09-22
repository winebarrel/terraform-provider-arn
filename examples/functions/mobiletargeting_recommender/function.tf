# arn:aws:mobiletargeting:ap-northeast-1:111111111111:recommenders/recommender-id
output "mobiletargeting_recommender" {
  value = provider::arn::mobiletargeting_recommender("recommender-id")
}
