# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/recommender-filters/recommender-filter-name
output "profile_recommender_filters" {
  value = provider::arn::profile_recommender_filters("domain-name", "recommender-filter-name")
}
