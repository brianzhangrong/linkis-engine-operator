package shellintercept

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func CreateShellEnv() {
	cmd := exec.Command("python")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("1111", err)
	}
	defer stdout.Close() // 保证关闭输出流

	cmd.Run()

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal("2222", err)
	} else {
		log.Println("3333", string(opBytes))
	}
}
