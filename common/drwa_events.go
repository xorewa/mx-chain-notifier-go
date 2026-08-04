package common

// DRWA event identifiers mirrored from mx-chain-es-indexer-go's canonical list.
// The notifier does not interpret payloads, but exposing the identifiers here
// lets subscribers build explicit filters without hard-coding strings.
const (
	DrwaAssetRegisteredEvent         = "drwaAssetRegistered"
	DrwaAssetUpdatedEvent            = "drwaAssetUpdated"
	DrwaTokenPolicyEvent             = "drwaTokenPolicy"
	DrwaHolderComplianceEvent        = "drwaHolderCompliance"
	DrwaTransferDeniedEvent          = "drwaTransferDenied"
	DrwaTransferAllowedEvent         = "drwaTransferAllowed"
	DrwaGlobalPauseEvent             = "drwaGlobalPause"
	DrwaMetadataProtectionEvent      = "drwaMetadataProtection"
	DrwaWhitePaperCidSetEvent        = "drwaWhitePaperCidSet"
	DrwaRegistrationStatusSetEvent   = "drwaRegistrationStatusSet"
	DrwaIdentityRegisteredEvent      = "drwaIdentityRegistered"
	DrwaComplianceUpdatedEvent       = "drwaComplianceUpdated"
	DrwaIdentityDeactivatedEvent     = "drwaIdentityDeactivated"
	DrwaIdentityErasedEvent          = "drwaIdentityErased"
	DrwaWindDownInitiatedEvent       = "drwaWindDownInitiated"
	DrwaAuditorProposedEvent         = "drwaAuditorProposed"
	DrwaAuditorAcceptedEvent         = "drwaAuditorAccepted"
	DrwaAuditorRevokedEvent          = "drwaAuditorRevoked"
	DrwaAttestationOverwrittenEvent  = "drwaAttestationOverwritten"
	DrwaAttestationRecordedEvent     = "drwaAttestationRecorded"
	DrwaGovernanceProposedEvent      = "drwaGovernanceProposed"
	DrwaGovernanceAcceptedEvent      = "drwaGovernanceAccepted"
	DrwaGovernanceRevokedEvent       = "drwaGovernanceRevoked"
	DrwaAuthActionProposedEvent      = "drwaAuthActionProposed"
	DrwaAuthActionSignedEvent        = "drwaAuthActionSigned"
	DrwaAuthActionUnsignedEvent      = "drwaAuthActionUnsigned"
	DrwaAuthActionDiscardedEvent     = "drwaAuthActionDiscarded"
	DrwaAuthActionPerformedEvent     = "drwaAuthActionPerformed"
	DrwaAuthorizedCallerUpdatedEvent = "drwaAuthorizedCallerUpdated"
	DrwaSignerAddedEvent             = "drwaSignerAdded"
	DrwaSignerRemovedEvent           = "drwaSignerRemoved"
	DrwaSignerReplacedEvent          = "drwaSignerReplaced"
	DrwaQuorumChangedEvent           = "drwaQuorumChanged"
)

// DRWAEventIdentifiers lists all supported DRWA event identifiers.
var DRWAEventIdentifiers = []string{
	DrwaAssetRegisteredEvent,
	DrwaAssetUpdatedEvent,
	DrwaTokenPolicyEvent,
	DrwaHolderComplianceEvent,
	DrwaTransferDeniedEvent,
	DrwaTransferAllowedEvent,
	DrwaGlobalPauseEvent,
	DrwaMetadataProtectionEvent,
	DrwaWhitePaperCidSetEvent,
	DrwaRegistrationStatusSetEvent,
	DrwaIdentityRegisteredEvent,
	DrwaComplianceUpdatedEvent,
	DrwaIdentityDeactivatedEvent,
	DrwaIdentityErasedEvent,
	DrwaWindDownInitiatedEvent,
	DrwaAuditorProposedEvent,
	DrwaAuditorAcceptedEvent,
	DrwaAuditorRevokedEvent,
	DrwaAttestationOverwrittenEvent,
	DrwaAttestationRecordedEvent,
	DrwaGovernanceProposedEvent,
	DrwaGovernanceAcceptedEvent,
	DrwaGovernanceRevokedEvent,
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
}

// IsDRWAIdentifier returns true if the identifier is one of the DRWA events.
func IsDRWAIdentifier(identifier string) bool {
	for _, item := range DRWAEventIdentifiers {
		if identifier == item {
			return true
		}
	}
	return false
}
