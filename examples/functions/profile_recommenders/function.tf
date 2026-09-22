# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/recommenders/recommender-type-name
output "profile_recommenders" {
  value = provider::arn::profile_recommenders("domain-name", "recommender-type-name")
}
