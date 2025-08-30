package enemy

//MenosGrandes factory itself should implement whole creation process, it does not need arguments
type EnemyFactoryI interface {
	CreateEnemy(ed EnemyData) EnemyI
}
