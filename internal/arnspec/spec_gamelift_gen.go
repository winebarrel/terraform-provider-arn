// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: gamelift
// Source: https://servicereference.us-east-1.amazonaws.com/v1/gamelift/gamelift.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "gamelift_alias", Service: "gamelift", Resource: "alias", Template: "arn:${Partition}:gamelift:${Region}::alias/${AliasId}"},
		{Name: "gamelift_build", Service: "gamelift", Resource: "build", Template: "arn:${Partition}:gamelift:${Region}:${Account}:build/${BuildId}"},
		{Name: "gamelift_container_fleet", Service: "gamelift", Resource: "containerFleet", Template: "arn:${Partition}:gamelift:${Region}:${Account}:containerfleet/${FleetId}"},
		{Name: "gamelift_container_group_definition", Service: "gamelift", Resource: "containerGroupDefinition", Template: "arn:${Partition}:gamelift:${Region}:${Account}:containergroupdefinition/${Name}"},
		{Name: "gamelift_fleet", Service: "gamelift", Resource: "fleet", Template: "arn:${Partition}:gamelift:${Region}:${Account}:fleet/${FleetId}"},
		{Name: "gamelift_game_server_group", Service: "gamelift", Resource: "gameServerGroup", Template: "arn:${Partition}:gamelift:${Region}:${Account}:gameservergroup/${GameServerGroupName}"},
		{Name: "gamelift_game_session_queue", Service: "gamelift", Resource: "gameSessionQueue", Template: "arn:${Partition}:gamelift:${Region}:${Account}:gamesessionqueue/${GameSessionQueueName}"},
		{Name: "gamelift_location", Service: "gamelift", Resource: "location", Template: "arn:${Partition}:gamelift:${Region}:${Account}:location/${LocationId}"},
		{Name: "gamelift_matchmaking_configuration", Service: "gamelift", Resource: "matchmakingConfiguration", Template: "arn:${Partition}:gamelift:${Region}:${Account}:matchmakingconfiguration/${MatchmakingConfigurationName}"},
		{Name: "gamelift_matchmaking_rule_set", Service: "gamelift", Resource: "matchmakingRuleSet", Template: "arn:${Partition}:gamelift:${Region}:${Account}:matchmakingruleset/${MatchmakingRuleSetName}"},
		{Name: "gamelift_script", Service: "gamelift", Resource: "script", Template: "arn:${Partition}:gamelift:${Region}:${Account}:script/${ScriptId}"},
	})
}
