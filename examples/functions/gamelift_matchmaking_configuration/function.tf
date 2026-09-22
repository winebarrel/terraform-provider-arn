# arn:aws:gamelift:ap-northeast-1:111111111111:matchmakingconfiguration/matchmaking-configuration-name
output "gamelift_matchmaking_configuration" {
  value = provider::arn::gamelift_matchmaking_configuration("matchmaking-configuration-name")
}
