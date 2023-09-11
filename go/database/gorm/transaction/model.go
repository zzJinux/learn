package transaction

type Alpha struct {
	ID  string
	Val string
}

type Beta struct {
	UID  string `gorm:"primaryKey"`
	Data string
}
