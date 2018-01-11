// invoke sys-mysql-diff scripts to monitor mysql diff
package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

var (
	commandPath string
	errNotFound = errors.New("sys-mysql-diff utility not found")
)

func commandCheck() error {
	if commandPath == "" {
		path, err := exec.LookPath("sys-mysql-diff")
		if err != nil {
			return errNotFound
		}
		commandPath = path
	}
	return nil
}

func startDiff(p *mysqlParams) (string, error) {
	if err := commandCheck(); err != nil {
		return "", err
	}
	//processInfo := fmt.Sprintf("-P %d", p.port)
	//if existsProcess(processInfo) {
	//	log.Printf("already exists mysql diff check process: host: %s, port: %d",
	//		p.host, p.port)
	//	return nil
	//}
	out, err := exec.Command(commandPath, "-h", p.host, "-P", strconv.FormatUint(uint64(p.port), 10),
		"-d", p.db, "-u", p.user, "-p", p.pass, "-t", "-r").CombinedOutput()
	if err != nil {
		fmt.Printf("execute command error: %s\n", string(out))
		return "", err
	}
	return string(out), nil
}

//func existsProcess(pInfo string) bool {
//	out , err := exec.Command("ps", "-ef").CombinedOutput()
//	if err != nil {
//		return false
//	}
//	return strings.Contains(string(out), pInfo)
//}
