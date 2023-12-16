package os_exec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestOsExec(t *testing.T) {
	//cmd := exec.Command("date", "-u '+%a, %d %b %Y %H:%M:%S GMT'")
	cmd := exec.Command("date", "-u", "+%a, %d %b %Y %H:%M:%S GMT")
	cmd.Env = []string{"LANG=en_US.UTF-8"}
	date, err := cmd.Output()
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(string(date))
}

func TestOsExecLs(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
}

// TestPipe 测试管道
func TestPipe(t *testing.T) {
	cmdCat := exec.Command("cat", "generics.go")
	catout, err := cmdCat.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get StdoutPipe of cat: %v", err)
	}
	cmdWC := exec.Command("wc", "-l")
	cmdWC.Stdin = catout
	cmdWC.Stdout = os.Stdout
	err = cmdCat.Start()
	if err != nil {
		log.Fatalf("failed to call cmdCat.Run(): %v", err)
	}
	err = cmdWC.Start()
	if err != nil {
		log.Fatalf("failed to call cmdWC.Start(): %v", err)
	}
	cmdCat.Wait()
	cmdWC.Wait()
}

func TestPipe2(t *testing.T) {
	cmdCat := exec.Command("cat", "exec_test.go")
	cmdWC := exec.Command("wc", "-l")
	data, err := pipeCommands(cmdCat, cmdWC)
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)

}
func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}

func TestBashPipe(t *testing.T) {
	cmd := exec.Command("bash", "-c", "cat exec_test.go| wc -l")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)
}
