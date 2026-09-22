# arn:aws:gamelift:ap-northeast-1:111111111111:matchmakingruleset/matchmaking-rule-set-name
output "gamelift_matchmaking_rule_set" {
  value = provider::arn::gamelift_matchmaking_rule_set("matchmaking-rule-set-name")
}
