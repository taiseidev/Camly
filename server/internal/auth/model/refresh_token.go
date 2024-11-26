package model

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`                                            // プライマリーキー
	UserID    uint      `gorm:"not null"`                                              // ユーザーID
	TokenHash string    `gorm:"size:255;not null"`                                     // トークンのハッシュ
	UserAgent string    `gorm:"size:255"`                                              // ユーザーエージェント
	IpAddress string    `gorm:"size:45"`                                               // IPアドレス
	ExpiresAt time.Time `gorm:"not null"`                                              // 期限
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`                             // 作成日時
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新日時
}

// 外部キー制約の追加
func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
