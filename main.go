/*
The MIT License (MIT)

mysqldiff - mysql diff tool
invoke sys-mysql-diff scripts
zhe.chen<chenzhe07@gmail.com>

*/

package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

var config struct {
	Conf      string
	Debug     bool
	Useremail string
}

func main() {
	logpath := flag.String("log", "", "")
	conf := flag.String("conf", "", "config file to verify database and firewall info")
	verbose := flag.Bool("verbose", false, "wheather print verbose message")

	flag.Parse()

	config.Conf = *conf

	if len(config.Conf) <= 0 {
		log.Printf("Error: You must secipfy the conf file")
		os.Exit(1)
	}
	if *logpath != "" {
		f, err := os.OpenFile(*logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		defer f.Close()

		log.SetOutput(f)
	}

	conf_fh, err := get_config(config.Conf)
	if err != nil {
		log.Printf("Error: read config file error")
		os.Exit(1)
	}

	backend_dsn, err := get_backend_dsn(conf_fh)
	backend_dbh, err := dbh(backend_dsn)
	if err != nil {
		log.Printf("get backend dbh error: %s", err)
		os.Exit(1)
	}

	log.Printf("begin ---------------------------")

	sections := get_mysql_list(conf_fh)
	for _, section := range sections {
		if !strings.EqualFold(section, "default") &&
			!strings.EqualFold(section, "backend") {
			//check mysql list
			startdiff := time.Now()
			mysqlval, err := get_mysql_instance(conf_fh, section)
			if err != nil {
				log.Printf("Warn: get instance for section [%s] error: %v\n", section, err)
				continue
			}
			mysqlval.changes, err = startDiff(&mysqlval)
			if err != nil {
				log.Printf("Warn: get changes for section [%s] error: %v\n", section, err)
				continue
			}
			elapsed := time.Since(startdiff)
			if *verbose {
				log.Printf("=> changes from %s:%d, elapsed %s\n%s", mysqlval.host, mysqlval.port, elapsed, mysqlval.changes)
			}
			if len(mysqlval.changes) == 0 {
				continue
			}
			if insertlog(backend_dbh, &mysqlval) {
				log.Printf("insert %s:%d/%s ok", mysqlval.host, mysqlval.port, mysqlval.db)
			}
		}
	}
	log.Printf("end ---------------------------")
}
