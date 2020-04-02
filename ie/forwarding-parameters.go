package ie

func NewForwardingParameters(teid uint16, destinationInterface *IE, outerHeaderCreation *IE) *IE {
	return newGroupedIE(ForwardingParameters, teid, destinationInterface, outerHeaderCreation)
}
