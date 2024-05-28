package model

func migration() {
	var err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&User{},
	)
	if err != nil {
		panic(err)
	} else {
		return
	}
}
