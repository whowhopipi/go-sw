//+build ignore

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func testIt(target string) (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command(path.Join(os.Getenv("PWD"), target))
	u, _ := user.Current()
	cmd.Env = append(cmd.Env, "USER="+u.Username)
	cmd.Env = append(cmd.Env, "HOME="+u.HomeDir)
	cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	cmd.Env = append(cmd.Env, "GOTRACEBACK=none")
	cmd.Env = append(cmd.Env, "PATH="+path.Join(
		os.Getenv("PWD"),
		"go-sw64", "bin", "linux_sw64")+":/usr/bin:/bin")
	cmd.Env = append(cmd.Env, "GOROOT="+path.Join(
		os.Getenv("PWD"), "go-sw64"))
	cmd.Dir = path.Dir(target)
	cmd.Stderr = &buf
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	done := make(chan bool)
	go func() {
		err = cmd.Wait()
		done <- true
	}()
	select {
	case <-time.After(Timeout):
		p := cmd.Process
		if p != nil {
			p.Kill()
		}
		return "", errors.New("timeout")
	case <-done:
		return buf.String(), err
	}
}

type Result struct {
	Name   string
	Output string
}

func fix_bin(root string) error {
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf("cp %s/bin/linux_sw64/* %s/bin", root, root))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

const Timeout = time.Second * 60 * 15
const MaxProcess = 14

var good []Result
var bad []Result

func main() {
	if len(os.Args) < 2 {
		println("Please set test target")
		return
	}

	root := os.Args[1]
	if fix_bin(root) != nil {
		println("Please move this program on parent of the go-sw64 directory.")
		return
	}

	all := make(chan string, MaxProcess)
	go func() {
		filepath.Walk(root, func(_path string, info os.FileInfo, err error) error {
			if strings.Contains(_path, "testdata/") {
				return nil
			}
			if strings.Contains(_path, "go/internal/load") {
				return nil
			}
			if !strings.HasSuffix(_path, ".test") {
				return nil
			}
			all <- _path
			return nil
		})
		close(all)
	}()
	RunAll(all)
}

func RunOne(t string) {
	begin := time.Now()
	output, err := testIt(t)
	duration := time.Now().Sub(begin)
	if len(output) > 3000 {
		output = output[:3000]
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "*****%s FAIL in %0.2fs*******\n", t, duration.Seconds())
		bad = append(bad, Result{
			Name:   t,
			Output: output + "\n" + err.Error(),
		})
	} else {
		fmt.Fprintf(os.Stderr, "%s PASS in %0.2fs\n", t, duration.Seconds())
		good = append(good, Result{
			Name:   t,
			Output: output,
		})
	}
}

func worker(jobs <-chan string) {
	for j := range jobs {
		RunOne(j)
	}
	allJobDone.Done()
}

var allJobDone sync.WaitGroup

func RunAll(targets chan string) {
	for j := 0; j < MaxProcess; j++ {
		allJobDone.Add(1)
		go worker(targets)
	}
	allJobDone.Wait()

	all := len(good) + len(bad)
	fmt.Printf("%d packages test at %v \n", all, time.Now())

	// fmt.Printf("Passe %d/%d\n========\n\n", len(good), all)
	// for i, r := range good {
	// 	fmt.Printf("%d %s\n", i+1, r.Name)
	// 	fmt.Println("---------------")
	// 	fmt.Println("```")
	// 	fmt.Println(r.Output)
	// 	fmt.Printf("```\n\n\n")
	// }

	fmt.Printf("Failed %d/%d\n========\n\n", len(bad), all)
	for i, r := range bad {
		fmt.Printf("%d %s\n", i+1, r.Name)
		fmt.Println("---------------")
		fmt.Println("```")
		fmt.Println(r.Output)
		fmt.Printf("```\n\n\n")
	}
}
