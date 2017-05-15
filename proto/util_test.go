package abaeve_auth

import "testing"

func TestAllianceAddTopic(t *testing.T) {
	topic := AllianceAddTopic()

	if topic != EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}

func TestAllianceDeleteTopic(t *testing.T) {
	topic := AllianceDeleteTopic()

	if topic != EntityType_ALLIANCE.String()+"."+EntityOperation_REMOVE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}

func TestCorporationAddTopic(t *testing.T) {
	topic := CorporationAddTopic()

	if topic != EntityType_CORPORATION.String()+"."+EntityOperation_ADD_OR_UPDATE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}

func TestCorporationDeleteTopic(t *testing.T) {
	topic := CorporationDeleteTopic()

	if topic != EntityType_CORPORATION.String()+"."+EntityOperation_REMOVE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}

func TestCharacterAddTopic(t *testing.T) {
	topic := CharacterAddTopic()

	if topic != EntityType_CHARACTER.String()+"."+EntityOperation_ADD_OR_UPDATE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}

func TestCharacterDeleteTopic(t *testing.T) {
	topic := CharacterDeleteTopic()

	if topic != EntityType_CHARACTER.String()+"."+EntityOperation_REMOVE.String() {
		t.Errorf("Expected: %s but received '%s'", EntityType_ALLIANCE.String()+"."+EntityOperation_ADD_OR_UPDATE.String(), topic)
	}
}
