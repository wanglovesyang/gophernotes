package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func eval_magic(codes []string) {
	tpLine := strings.Trim(codes[0], "%")
	if strings.HasPrefix(tpLine, "bash") {
		bashCode := strings.Join(codes[1:], "\n")
		cmd := exec.Command("sh", "-c", bashCode)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else if strings.HasPrefix(tpLine, "put") {
		args := strings.Split(tpLine, "put")
		if len(args) <= 1 {
			return
		}
		cmd := fmt.Sprintf("hadoop fs -put %s hdfs://172.29.15.131:8020/h2o/%s", args[1], os.Getenv("USER"))
		c := exec.Command(cmd)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	}

}
