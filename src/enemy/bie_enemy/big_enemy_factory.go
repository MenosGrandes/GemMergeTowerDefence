package big_enemy

import "example/hello/src/enemy"

type BigEnemyFactory struct {
}

func (ef *BigEnemyFactory) CreateEnemy(ed enemy.EnemyData) enemy.EnemyI {
	return &BigEnemy{enemyData: ed}
}
