package common

import "testing"

func TestDRWAEventIdentifiers_NonEmpty(t *testing.T) {
	for _, id := range DRWAEventIdentifiers {
		if id == "" {
			t.Fatalf("DRWA event identifier must not be empty")
		}
	}
}

func TestDRWAEventIdentifiers_NoDuplicates(t *testing.T) {
	seen := make(map[string]struct{}, len(DRWAEventIdentifiers))
	for _, id := range DRWAEventIdentifiers {
		if _, exists := seen[id]; exists {
			t.Fatalf("duplicate DRWA event identifier: %q", id)
		}
		seen[id] = struct{}{}
	}
}

func TestIsDRWAIdentifier(t *testing.T) {
	if !IsDRWAIdentifier(DrwaTokenPolicyEvent) {
		t.Fatalf("expected %q to be a DRWA identifier", DrwaTokenPolicyEvent)
	}
	if IsDRWAIdentifier("nonDrwaEvent") {
		t.Fatalf("unexpected DRWA identifier match for non DRWA event")
	}
}

func TestIsDRWAIdentifier_AllCanonicalEvents(t *testing.T) {
	for _, identifier := range DRWAEventIdentifiers {
		if !IsDRWAIdentifier(identifier) {
			t.Fatalf("expected %q to be recognized as a DRWA identifier", identifier)
		}
	}
}

func TestIsDRWAIdentifier_PreviouslyOmittedIndexerEvents(t *testing.T) {
	for _, identifier := range []string{
		DrwaAssetUpdatedEvent,
		DrwaAuthActionProposedEvent,
		DrwaAuthActionSignedEvent,
		DrwaAuthActionUnsignedEvent,
		DrwaAuthActionDiscardedEvent,
		DrwaAuthActionPerformedEvent,
		DrwaAuthorizedCallerUpdatedEvent,
		DrwaSignerAddedEvent,
		DrwaSignerRemovedEvent,
		DrwaSignerReplacedEvent,
		DrwaQuorumChangedEvent,
	} {
		if !IsDRWAIdentifier(identifier) {
			t.Fatalf("expected previously omitted event %q to be recognized", identifier)
		}
	}
}

// Keep this reviewable list in sync with the indexer's canonical allow-list.
// It is intentionally string-based so renaming a constant cannot hide drift.
func TestDRWAEventIdentifiers_IndexerParity(t *testing.T) {
	expected := []string{
		"drwaAssetRegistered", "drwaAssetUpdated", "drwaTokenPolicy", "drwaHolderCompliance",
		"drwaTransferDenied", "drwaTransferAllowed", "drwaGlobalPause", "drwaMetadataProtection",
		"drwaWhitePaperCidSet", "drwaRegistrationStatusSet", "drwaIdentityRegistered",
		"drwaComplianceUpdated", "drwaIdentityDeactivated", "drwaIdentityErased", "drwaWindDownInitiated",
		"drwaAuditorProposed", "drwaAuditorAccepted", "drwaAuditorRevoked", "drwaAttestationOverwritten",
		"drwaAttestationRecorded", "drwaGovernanceProposed", "drwaGovernanceAccepted", "drwaGovernanceRevoked",
		"drwaAuthActionProposed", "drwaAuthActionSigned", "drwaAuthActionUnsigned", "drwaAuthActionDiscarded",
		"drwaAuthActionPerformed", "drwaAuthorizedCallerUpdated", "drwaSignerAdded", "drwaSignerRemoved",
		"drwaSignerReplaced", "drwaQuorumChanged",
	}
	if len(DRWAEventIdentifiers) != len(expected) {
		t.Fatalf("DRWA registry length %d differs from canonical length %d", len(DRWAEventIdentifiers), len(expected))
	}
	for _, identifier := range expected {
		if !IsDRWAIdentifier(identifier) {
			t.Fatalf("DRWA registry is missing canonical event %q", identifier)
		}
	}
}
