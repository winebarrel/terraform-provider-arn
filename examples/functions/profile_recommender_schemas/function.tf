# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/recommender-schemas/recommender-schema-name
output "profile_recommender_schemas" {
  value = provider::arn::profile_recommender_schemas("domain-name", "recommender-schema-name")
}
