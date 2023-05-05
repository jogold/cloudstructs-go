// High-level constructs for AWS CDK
package cloudstructs


// The DMARC alignment mode.
type DmarcAlignment string

const (
	// Relaxed alignment mode.
	DmarcAlignment_RELAXED DmarcAlignment = "RELAXED"
	// Strict alignment mode.
	DmarcAlignment_STRICT DmarcAlignment = "STRICT"
)

