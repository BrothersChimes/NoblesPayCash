package item

// Type designates information for items that can be bought and sold
type Type struct {
	Name       string
	PluralName string
	UseAn      bool
	IsWeapon   bool
}
