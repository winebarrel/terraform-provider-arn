# arn:aws:gamelift:ap-northeast-1:111111111111:gamesessionqueue/game-session-queue-name
output "gamelift_game_session_queue" {
  value = provider::arn::gamelift_game_session_queue("game-session-queue-name")
}
