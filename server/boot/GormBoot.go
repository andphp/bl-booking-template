package boot

import (
	"fmt"
	"time"

	"github.com/andphp/go-gin/goby"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormBoot() {
	go func() {
		if err := GormMysql(); err != nil {
			// FatalErrorChan <- err
			panic(fmt.Errorf("mysql connect error:%s", err))
		}
	}()
}

func GormMysql() error {
	m := goby.GOBY_CONFIG.Mysql
	if m.Dbname == "" {
		panic("dbname is null")
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(m.MaxLifetime))
		goby.GOBY_DB = db
		return nil
	}
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch goby.GOBY_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = goby.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = goby.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = goby.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = goby.Default.LogMode(logger.Info)
	default:
		config.Logger = goby.Default.LogMode(logger.Info)
	}
	return config
}
