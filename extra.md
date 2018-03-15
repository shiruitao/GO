```go
type ExtraInfo struct {
	ID     uint32 `gorm:"column:id";primary_key;auto_increment" json:"id"`
	UserID uint32 `gorm:"column:user_id;not null;";json:"user_id"`
	Field  string `gorm:"column:field"json:"field"`
	Value  string `gorm:"column:value"json:"value"`
	Status uint8  `gorm:"column:status;not null;";json:"status"`
}

ValueInfo struct {
		UserID uint32 `json:"user_id"`
		Field  string `json:"field"`
		Value  string `json:"value"`
	}

// bbs: add field and value
func (this *UserServiceProvider) PutExtraValue(conn orm.Connection, info *ValueInfo) error {
	var (
		extraInfo ExtraInfo
	)

	extraInfo.UserID = info.UserID
	extraInfo.Field = info.Field
	extraInfo.Value = info.Value

	db := conn.(*gorm.DB)

	return db.Create(&extraInfo).Error
}

// Change ExtraInfo information
func (this *UserServiceProvider) ChangeExtraInfo(conn orm.Connection, info *ValueInfo) error {
	var (
		extraInfo ExtraInfo
	)

	db := conn.(*gorm.DB)
	return db.Model(&extraInfo).Where("user_id = ? AND field = ? AND status = ?", info.UserID, info.Field, 0).Update("value", info.Value).Error
}

// bbs: get user information by userId.
func (this *UserServiceProvider) GetExtraInfo(conn orm.Connection, userID uint32) (*[]ExtraInfo, error) {
	var (
		extraInfo []ExtraInfo
	)

	db := conn.(*gorm.DB)
	err := db.Where("user_id = ?", userID).Find(&extraInfo).Error

	return &extraInfo, err
}

// bbs: get user value by userId and field.
func (this *UserServiceProvider) GetExtraValue(conn orm.Connection, userID uint32, field string) (*string, error) {
	var (
		extraInfo ExtraInfo
	)
	db := conn.(*gorm.DB)
	err := db.Where("user_id = ? AND field = ? ", userID, field).Find(&extraInfo).Error

	return &extraInfo.Value, err
}

// bbs: modify status
// 0 -> normal
// 1 -> delete or hide
func (this *UserServiceProvider) ChangeExtraStatus(conn orm.Connection, id uint32, status uint8) error {
	var (
		extraInfo ExtraInfo
	)

	db := conn.(*gorm.DB)
	extraInfo.Status = status

	return db.Model(&extraInfo).Where("id = ?", id).Update("status", status).Error
}
```

CREATE TABLE IF NOT EXISTS `extra_infos` (
  `id` INT(32) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT(32) UNSIGNED NOT NULL,
  `field` VARCHAR(256),
  `value` VARCHAR(256),
  `status` INT(8) UNSIGNED,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
