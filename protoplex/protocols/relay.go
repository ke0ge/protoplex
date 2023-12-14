package protocols

// NewRelayProtocol initializes any Protocol to relay.
func NewRelayProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:   "Relay",
		Target: targetAddress,
	}
}
