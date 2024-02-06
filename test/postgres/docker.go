package postgres

import (
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	database = "test"
	password = "admin"
)

func DbConnectionInit() (*gorm.DB, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	runDockerOpt := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env:        []string{"POSTGRES_PASSWORD=" + password, "POSTGRES_DB=" + database},
	}

	dockerConf := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.NeverRestart()
	}

	resource, err := pool.RunWithOptions(runDockerOpt, dockerConf)
	if err != nil {
		panic(err)
	}

	fnCleanup := func() {
		err := resource.Close()
		if err != nil {
			panic(err)
		}
	}

	connStr := fmt.Sprintf("host=localhost port=%s user=postgres dbname=%s password =%s sslmode=disable",
		resource.GetPort("5432/tcp"),
		database,
		password,
	)

	var gdb *gorm.DB

	err = pool.Retry(func() error {
		gdb, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			return err
		}
		db, err := gdb.DB()
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		panic(err)
	}

	return gdb, fnCleanup
}
