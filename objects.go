package nlib


type DB struct {
	DbType			DBType
	DSN					string
}

/*
type JoinType struct {
	Source   JoinSourceType
	Method   JoinMethodType
}

type PostItem struct {
	ID            libuuid.UUID        `json:"uuid" gorm:"column:id;primary_key;not null;default:uuid_generate_v4()"`
	RawText     	string        		  `json:"raw_text"`
	Name   				string
	FileSize      int64            		`json:"file_size" gorm:"type:bigint;not null;default:0"`
	CreatedAt     time.Time           `json:"created_at" gorm:"default:now()"`
}

*/
