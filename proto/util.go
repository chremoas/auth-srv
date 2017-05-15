package abaeve_auth

func AllianceAddTopic() string {
	return EntityType_ALLIANCE.String() + "." + EntityOperation_ADD_OR_UPDATE.String()
}

func AllianceDeleteTopic() string {
	return EntityType_ALLIANCE.String() + "." + EntityOperation_REMOVE.String()
}

func CorporationAddTopic() string {
	return EntityType_CORPORATION.String() + "." + EntityOperation_ADD_OR_UPDATE.String()
}

func CorporationDeleteTopic() string {
	return EntityType_CORPORATION.String() + "." + EntityOperation_REMOVE.String()
}

func CharacterAddTopic() string {
	return EntityType_CHARACTER.String() + "." + EntityOperation_ADD_OR_UPDATE.String()
}

func CharacterDeleteTopic() string {
	return EntityType_CHARACTER.String() + "." + EntityOperation_REMOVE.String()
}
