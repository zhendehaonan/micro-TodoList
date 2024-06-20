package model

func migration() {
	_ = _db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
}
