package ie

func NewPDI(teid uint16, sourceInterface *IE, fteid *IE, ueIpAddress *IE) *IE {
	return newGroupedIE(PDI, teid, sourceInterface, fteid, ueIpAddress)
}
