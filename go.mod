module github.com/k1LoW/tbls

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobuffalo/packr v1.24.1
	github.com/labstack/gommon v0.2.8
	github.com/lib/pq v1.0.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mattn/go-runewidth v0.0.4
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/spf13/cobra v0.0.3
	github.com/xo/dburl v0.0.0-20190203050942-98997a05b24f
	golang.org/x/sys v0.0.0-20190405154228-4b34438f7a67 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/watarukura/tbls/drivers/redshift => drivers/redshift
