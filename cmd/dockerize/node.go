package dockerize

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	projectName string
	postgres    bool
	redis       bool
	// port        string
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "dockerize node apps",
	Long:  `command pallets to generate nodejs app boilerplate with docker configured`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		list := []string{
			"/setup.sh",
			"/app/Dockerfile",
			"/app/Dockerfile.dev",
		}
		fmt.Print("What is your app name? ")
		projectName, _ = reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)

		fmt.Print("Do you want to enable Postgres? (y/n): ")
		postgresResponse, _ := reader.ReadString('\n')
		postgres = strings.ToLower(strings.TrimSpace(postgresResponse)) == "y"

		fmt.Print("Do you want to enable Redis? (y/n): ")
		redisResponse, _ := reader.ReadString('\n')
		redis = strings.ToLower(strings.TrimSpace(redisResponse)) == "y"

		// fmt.Print("Port number for your node app e.g. 7000 ")
		// portResponse, _ := reader.ReadString('\n')
		// port = strings.TrimSpace(portResponse)

		postgresArg := "n"
		redisArg := "n"
		if postgres {
			fmt.Printf("Postgres enabled: %t\n", postgres)
			postgresArg = "y"
			list = append(list, "/postgres/Dockerfile", "/db/postgres.js")
		}
		fmt.Printf("Postgres arg: %s\n", postgresArg)

		if redis {
			fmt.Printf("Redis enabled: %t\n", redis)
			redisArg = "y"
			list = append(list, "/db/redis.js")
		}
		fmt.Printf("Redis arg: %s\n", redisArg)

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		projectPath := filepath.Join(cwd, projectName)
		scriptPath := filepath.Join(cwd, "setup.sh")
		err = downloadFiles(list, cwd)
		if err != nil {
			fmt.Println(err)
			return
		}

		command := exec.Command("bash", scriptPath, projectPath, projectName, postgresArg, redisArg)
		err = command.Run()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func downloadFiles(paths []string, cwd string) error {
	for _, path := range paths {
		dir := filepath.Dir(path)

		dirPath := filepath.Join(cwd, dir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
		filePath := filepath.Join(dirPath, filepath.Base(path))

		// url := fmt.Sprintf("http://ubuntu:7000%s", path)
		url := fmt.Sprintf("https://raw.githubusercontent.com/qntum-dev/templates/main/node-docker%s", path)

		err = downloadFile(url, filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func init() {
	DockerizeCmd.AddCommand(nodeCmd)
}
