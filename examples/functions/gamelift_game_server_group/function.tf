# arn:aws:gamelift:ap-northeast-1:111111111111:gameservergroup/game-server-group-name
output "gamelift_game_server_group" {
  value = provider::arn::gamelift_game_server_group("game-server-group-name")
}
