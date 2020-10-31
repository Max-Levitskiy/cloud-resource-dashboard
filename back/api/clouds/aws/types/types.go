package types

type RegionalResourceScanner interface {
	Scan(accountId *string, region string, profileName string)
}
type GlobalResourceScanner interface {
	Scan(accountId *string, profileName *string)
}
