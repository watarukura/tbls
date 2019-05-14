package datasource

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/k1LoW/tbls/drivers"
	"github.com/k1LoW/tbls/drivers/mysql"
	"github.com/k1LoW/tbls/drivers/postgres"
	"github.com/k1LoW/tbls/drivers/sqlite"
	"github.com/k1LoW/tbls/schema"
	"github.com/pkg/errors"
	"github.com/watarukura/tbls/drivers/redshift"
	"github.com/xo/dburl"
)

// Analyze database
func Analyze(urlstr string) (*schema.Schema, error) {
	if strings.Index(urlstr, "json://") == 0 {
		return AnalizeJSON(urlstr)
	}
	s := &schema.Schema{}
	u, err := dburl.Parse(urlstr)
	if err != nil {
		return s, errors.WithStack(err)
	}
	splitted := strings.Split(u.Short(), "/")
	if len(splitted) < 2 {
		return s, errors.WithStack(fmt.Errorf("invalid DSN: parse %s -> %#v", urlstr, u))
	}

	db, err := dburl.Open(urlstr)
	defer db.Close()
	if err != nil {
		return s, errors.WithStack(err)
	}
	if err = db.Ping(); err != nil {
		return s, errors.WithStack(err)
	}

	var driver drivers.Driver

	switch u.Driver {
	case "postgres":
		s.Name = splitted[1]
		driver = new(postgres.Postgres)
	case "mysql":
		s.Name = splitted[1]
		driver = new(mysql.Mysql)
	case "sqlite3":
		s.Name = splitted[len(splitted)-1]
		driver = new(sqlite.Sqlite)
	case "redshift":
		s.Name = splitted[len(splitted)-1]
		driver = new(redshift.Redshift)
	default:
		return s, errors.WithStack(fmt.Errorf("unsupported driver '%s'", u.Driver))
	}
	d, err := driver.Info(db)
	if err != nil {
		return s, err
	}
	s.Driver = d
	err = driver.Analyze(db, s)
	if err != nil {
		return s, err
	}
	return s, nil
}

// AnalizeJSON analize `json://`
func AnalizeJSON(urlstr string) (*schema.Schema, error) {
	s := &schema.Schema{}
	splitted := strings.Split(urlstr, "json://")
	file, err := os.Open(splitted[1])
	if err != nil {
		return s, errors.WithStack(err)
	}
	dec := json.NewDecoder(file)
	dec.Decode(s)
	err = s.Repair()
	if err != nil {
		return s, errors.WithStack(err)
	}
	return s, nil
}
