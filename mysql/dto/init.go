package database

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/lworkltd/kits/service/restful/code"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	db                *gorm.DB
	log               = logrus.WithField("pkg", "dto")
	initOnce          sync.Once
	constMaxIdleConns = 20  // 单proxy最大空闲连接数
	constMaxOpenConns = 100 // 单proxy最大连接数
)

// DB wrapper if user want using readdb
func DB() *gorm.DB {
	return db
}

// Init 初始化
func Init() error {
	var err error
	initOnce.Do(func() {
		since := time.Now()
		//本地测试
		// mysqlUrl := appconf.GetApp().MysqlUrl
		mysqlUrl := "root:1234567@tcp(localhost:3306)/cryptobroker_9001?charset=utf8mb4&parseTime=True&loc=Local"
		endpoint := fmt.Sprintf("%s", mysqlUrl)
		err = InitMysql(&db, endpoint)
		if err != nil {
			panic(fmt.Sprintf("init mysql failed %v", err))
		}
		log.Infof("init mysql  done")
		log.Infof("init mysql cost %v", time.Since(since).Round(time.Millisecond))
	})

	return err
}

func InitMysql(db **gorm.DB, url string) error {

	d, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}
	d.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(url)},
		// Replicas: []gorm.Dialector{mysql.Open(urlReplica)},
		Policy: dbresolver.RandomPolicy{},
	}).
		SetMaxIdleConns(3).
		SetMaxOpenConns(50).
		SetConnMaxLifetime(time.Hour * 7),
	)
	log.WithFields(logrus.Fields{
		"constMaxIdleConns": constMaxIdleConns,
		"constMaxOpenConns": constMaxOpenConns,
	}).Info("conns")

	*db = d

	log.Print("Init MySQL")

	return nil
}

// dbError2CodeError 将DB-Driver产生得错误转化为通用错误
// 错误码详细参考：https://dev.mysql.com/doc/refman/5.5/en/server-error-reference.html
// 请按需添加。
func dbError2CodeError(err error) code.Error {
	if err.Error() == "record not found" {
		return code.NewMcodef("NOT_FOUND", err.Error())
	}

	// if err == mgo.ErrNotFound {
	// 	return code.New(int(tradepb.ErrorCode_GMT_COMM_NOTFOUND), err.Error())
	// }

	if strings.Index(err.Error(), "Error 1062") != -1 {
		return code.NewMcodef("DUPLICATED", err.Error())
	}

	return code.NewMcodef("DB_ERROR", err.Error())
}

// IsNotFound 判断是否是一个未查到目标
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	cerr, ok := err.(code.Error)
	if !ok {
		return false
	}

	if cerr.Mcode() == "NOT_FOUND" {
		return true
	}

	return false
}

// IsDuplicated .
func IsDuplicated(err error) bool {
	if err == nil {
		return false
	}
	cerr, ok := err.(code.Error)
	if !ok {
		return false
	}

	if cerr.Mcode() == "DUPLICATED" {
		return true
	}

	return false
}

func isDBDuplicated(err error) bool {
	return nil != err && strings.Contains(err.Error(), "Error 1062")
}

func IsDBNotFound(dbError error) bool {
	return dbError.Error() == "record not found"
}

func initTables(exchange string, tables ...interface{}) {
	db := DB()
	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			continue
		}
		db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(table)
	}
}

// TranscationItem 事务操作
type TranscationItem interface {
	Save(*gorm.DB) error
}

// DoTranscations 将savers作为一个事务整体处理
func DoTranscations(exchange string, savers []TranscationItem) error {
	// 单个不参与事务
	if len(savers) == 1 {
		return savers[0].Save(nil)
	}

	db := DB()

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	for _, saver := range savers {
		if err := saver.Save(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
