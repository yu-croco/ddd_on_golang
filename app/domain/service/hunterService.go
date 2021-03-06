package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateAttackMonsterDamage(hunter *model.Hunter, monster *model.Monster) model.HunterOffensePower {
	if monster.DefencePower.BiggerOrSameThan(hunter.OffensePower) {
		return hunter.OffensePower
	} else {
		return hunter.OffensePower.Twice().Minus(monster.DefencePower)
	}
}
