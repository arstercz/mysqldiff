/*config read to verify normal user*/
package main

import (
	"github.com/chenzhe07/goconfig"
)

func get_config(conf string) (c *goconfig.ConfigFile, err error) {
	c, err = goconfig.ReadConfigFile(conf)
	if err != nil {
		return c, err
	}
	return c, nil
}

func get_backend_dsn(c *goconfig.ConfigFile) (dsn string, err error) {
	dsn, err = c.GetString("backend", "dsn")
	if err != nil {
		return dsn, err
	}
	return dsn, nil
}

func get_mysql_instance(c *goconfig.ConfigFile, instance string) (t mysqlParams, err error) {
	t.host, err = c.GetString(instance, "host")
	t.port, err = c.GetInt64(instance, "port")
	t.db,   err = c.GetString(instance, "db")
	t.user, err = c.GetString(instance, "user")
	t.pass, err = c.GetString(instance, "pass")
	t.tag,  err = c.GetString(instance, "tag")

	if err != nil {
		return t, err
	}
	return t, nil
}

func get_mysql_list(c *goconfig.ConfigFile) (sections []string) {
	sections = c.GetSections()
	return sections
}
