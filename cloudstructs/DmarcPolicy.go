package cloudstructs


// The DMARC policy to apply to messages that fail DMARC compliance.
type DmarcPolicy string

const (
	// Do not apply any special handling to messages that fail DMARC compliance.
	DmarcPolicy_NONE DmarcPolicy = "NONE"
	// Quarantine messages that fail DMARC compliance.
	//
	// (usually by sending them to spam).
	DmarcPolicy_QUARANTINE DmarcPolicy = "QUARANTINE"
	// Reject messages that fail DMARC compliance.
	//
	// (usually by rejecting them outright).
	DmarcPolicy_REJECT DmarcPolicy = "REJECT"
)

