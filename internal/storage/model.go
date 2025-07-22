package storage

import "time"

// Модель запроса и XML-ответа
type Request struct {
	ID         uint `gorm:"primaryKey"`
	DateParam  string
	XMLContent string
	CreatedAt  time.Time
}
