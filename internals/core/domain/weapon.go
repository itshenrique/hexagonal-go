package domain

type Weapon struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AttackValue uint32 `json:"attackValue"`
}
