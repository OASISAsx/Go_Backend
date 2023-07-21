package infrastructure

import (
	"collection/internal/core/port"
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		DryRun: false,
		Logger: &SqlLogger{},
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	db.AutoMigrate(&port.Role{}, &port.Review{}, &port.Product{}, &port.Buycs{}, &port.BuyDetail{}, &port.Areapv{},
		&port.UserDetail{}, &port.Areaap{}, &port.Areadt{}, &port.Register{} ,&port.Cart{}  ) 

}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n", sql)
}
