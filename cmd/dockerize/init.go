package dockerize

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"path/filepath"

// 	"github.com/spf13/cobra"
// )

// var (
// 	projectName string
// 	postgres    bool
// 	redis       bool
// 	port        string
// )

// // dockerizeCmd represents the dockerize command
// var initCmd = &cobra.Command{
// 	Use:   "init",
// 	Short: "dockerize node apps",
// 	Long:  `command pallets to generate nodejs app boilerplate with docker configured`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		postgresArg := "n"
// 		redisArg := "n"
// 		if postgres {
// 			fmt.Printf("Postgres enabled %t\n", postgres)
// 			postgresArg = "y"
// 		}
// 		fmt.Printf("Postgres arg %s\n", postgresArg)

// 		if redis {
// 			fmt.Printf("Redis enabled %t\n", redis)

// 			redisArg = "y"
// 		}

// 		fmt.Printf("Redis arg %s\n", redisArg)

// 		cwd, err := os.Getwd()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		list := []string{
// 			"/setup.sh",
// 			"/app/Dockerfile",
// 			"/app/Dockerfile.dev",
// 			"/docker-compose.yml",
// 			"/docker-compose.dev.yml",
// 			"/postgres/Dockerfile",
// 			"/db/postgres.js",
// 			"/db/redis.js",
// 		}

// 		// fmt.Println(list)
// 		projectPath := filepath.Join(cwd, projectName)
// 		scriptPath := filepath.Join(cwd, "setup.sh")
// 		err = downloadFiles(list, cwd)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		command := exec.Command("bash", scriptPath, projectPath, projectName, postgresArg, redisArg)
// 		err = command.Run()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	},
// }

// func downloadFiles(paths []string, cwd string) error {
// 	for _, path := range paths {
// 		dir := filepath.Dir(path)

// 		dirPath := filepath.Join(cwd, dir)
// 		err := os.MkdirAll(dirPath, 0755)
// 		if err != nil {
// 			return err
// 		}
// 		// fmt.Println(dirPath)
// 		filePath := filepath.Join(dirPath, filepath.Base(path))

// 		// headUrl:=""

// 		// url := fmt.Sprintf("https://raw.githubusercontent.com/qntum-dev/templates/main/node-docker%s", path)

// 		url := fmt.Sprintf("http://ubuntu:7000%s", path)

// 		err = downloadFile(url, filePath)
// 		if err != nil {
// 			return err
// 		}

// 	}
// 	return nil
// }

// func downloadFile(url, filepath string) error {
// 	// fmt.Println(url)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }

// func init() {
// 	initCmd.Flags().StringVarP(&port, "port", "p", "7000", "mention the nodejs app port")
// 	initCmd.PersistentFlags().StringVarP(&projectName, "appname", "n", "app", "mention the node js app name")
// 	initCmd.Flags().BoolVarP(&postgres, "postgres", "d", false, "Enable Postgres")
// 	initCmd.Flags().BoolVarP(&redis, "redis", "r", false, "Enable redis")
// 	DockerizeCmd.AddCommand(initCmd)
// }
