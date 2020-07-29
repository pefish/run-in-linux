package main

import (
	"flag"
	"fmt"
	"github.com/pefish/run-in-linux/shell"
	"log"
	"os"
	"strings"
)

// docker run -t -i -v `pwd`:/app -w /app -e GOPROXY=https://goproxy.io golang:1.14 go run ./bin/test

func main() {
	flagSet := flag.NewFlagSet("run-in-linux", flag.ExitOnError)
	env := flagSet.String("e", "GOPROXY=https://goproxy.io", "环境变量，可以多次指定。例如 GOPROXY=https://goproxy.io,ABC=abc")
	port := flagSet.String("p", "", "端口映射。例如 3333:80,3334=81")
	image := flagSet.String("i", "golang:1.14", "镜像名")
	flagSet.Usage = func() {
		fmt.Fprintf(flagSet.Output(), "\n%s%s 是一个切换环境执行命令的工具\n\n", strings.ToUpper(string(flagSet.Name()[0])), flagSet.Name()[1:])
		fmt.Fprintf(flagSet.Output(), "Usage of %s:\n", flagSet.Name())
		flagSet.PrintDefaults()
		fmt.Fprintf(flagSet.Output(), "\n")
	}
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	args := flagSet.Args()
	if len(args) == 0 {
		log.Fatal("请指定要执行的命令")
	}

	cmdArr := make([]string, 0, 10)
	cmdArr = append(cmdArr, "docker", "run", "-t", "-i", "-v", "`pwd`:/app", "-w", "/app")
	for _, e := range strings.Split(*env, ",") {
		cmdArr = append(cmdArr, "-e", e)
	}
	if port != nil {
		for _, p := range strings.Split(*port, ",") {
			cmdArr = append(cmdArr, "-p", p)
		}
	}
	cmdArr = append(cmdArr, *image, strings.Join(args, " "))

	cmd, err := shell.GetCmd(strings.Join(cmdArr, " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.String())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	for _, e := range os.Environ() {
		cmd.Env = append(cmd.Env, e)
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
