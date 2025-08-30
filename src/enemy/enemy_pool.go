package enemy

// var instance map[EnemyType]*pool.FixedPool[EnemyI]
// var once sync.Once

// func GetEnemyPool(et EnemyType) *pool.FixedPool[EnemyI] {
// 	once.Do(func() {
// 		instance = pool.NewFixedPool(func(i constants.ID) EnemyI {
// 			return *NewEnemyPtr(movable.NewMovableObjectEmpty(i), drawable.NewDrawableEmptyPtr())
// 		})
// 	})
// 	return instance
// }
