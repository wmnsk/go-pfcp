package ie

func NewCreatePDR(teid uint16, pdrid *IE, precedence *IE, pdi *IE, outerHeaderRemoval *IE, farID *IE) *IE {
	return newGroupedIE(CreatePDR, teid, pdrid, precedence, pdi, outerHeaderRemoval, farID)
}
