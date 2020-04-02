package ie

func NewCreateFAR(teid uint16, farID *IE, applyAction *IE, forwardingParameter *IE) *IE {
	return newGroupedIE(CreateFAR, teid, farID, applyAction, forwardingParameter)
}
