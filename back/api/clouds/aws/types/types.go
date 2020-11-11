package types

type RegionalResourceScanner interface {
	Scan(projectId *string, region string, profileName string)
}
type GlobalResourceScanner interface {
	Scan(projectId *string, profileName *string)
}
