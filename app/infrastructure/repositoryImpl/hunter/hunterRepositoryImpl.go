package hunter

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model/hunter"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

// ToDo: 循環参照をどうにかする
func FindById(db *gorm.DB, id hunter.Id) {
	var hunterDao *dao.Hunter
	db.First(&hunterDao, int(id))

	//return hunterDao.ConvertToModel()
}

// ToDo: 循環参照をどうにかする
func Update(db *gorm.DB, hunter dao.Hunter) {
	var hunterDao *dao.Hunter

	db.First(&hunterDao, int(hunter.ID))
	hunterDao.Life = hunter.Life

	db.Save(&hunterDao)
	//return hunterDao.ConvertToModel()
}
