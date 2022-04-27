package data_test

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"testing"
	"user/internal/conf"
	"user/internal/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "test biz data ")
}

var cleaner func()
var Db *data.Data
var ctx context.Context

func initialize(db *gorm.DB) error {
	err := db.AutoMigrate(
		&data.User{},
	)
	return errors.WithStack(err)
}

var _ = BeforeSuite(func() {
	con, f := data.DockerMysql("mariadb", "latest")
	cleaner = f
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: con}}
	db := data.NewDB(config)
	mysqlDb, _, err := data.NewData(config, nil, db, nil)
	if err != nil {
		return
	}
	Db = mysqlDb
	err = initialize(db)
	if err != nil {
		return
	}
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	cleaner()
})
