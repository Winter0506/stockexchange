package model

import (
	"database/sql/driver"
	"encoding/json"
)

// have use model.GoodsDetail
type StockDetail struct {
	Stock int32
	Num   int32
}

func (g StockDetail) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *StockDetail) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

// have use  model.Inventory
type Inventory struct {
	BaseModel
	Stock int32 `gorm:"type:int;index"`
	Total int32 `gorm:"type:int"`
	// Version int32 `gorm:"type:int"` //分布式锁的乐观锁  这个东西在做乐观锁时候有用  现在已经没有用了
}

// have use model.StockSellDetail
type StockSellDetail struct {
	TrustSn string      `gorm:"type:varchar(30);index:idx_order_sn,unique;"`
	Status  int32       `gorm:"type:varchar(10)"` //1 表示已扣减 2. 表示已归还
	Detail  StockDetail `gorm:"type:varchar(200)"`
}

func (Inventory) TableName() string {
	return "inventory"
}
