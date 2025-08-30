package atack_data

type AttackData struct {
	Health int64
	Armor  int64
}

func NewAttackData(Health, Armor int64) AttackData {
	return AttackData{Health: Health, Armor: Armor}
}
