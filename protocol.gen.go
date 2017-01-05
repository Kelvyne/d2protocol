package d2protocol

type Version struct {
	Major    uint
	Minor    uint
	Release  uint
	Revision uint
	Patch    uint
}

// Version is the current version of the Dofus 2 protocol
var ProtocolVersion = Version{2, 39, 0, 117122, 0}
