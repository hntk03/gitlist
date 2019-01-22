
package main

import (
	"fmt"
	"os"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"io/ioutil"
)

func run() int{


    f, err := os.Open("~/gitlist")
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()

    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)
    // 出力
    fmt.Println(string(b))



	return 0

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil{
		fmt.Fprintf(os.Stderr, "[ERROR] couldnt acquire\n")
		return 1
	}

	// fmt.Println(repos["CloneURL"])
	// fmt.Println(reflect.TypeOf(repos))
	for i := 0; i < len(repos) ; i++{
		fmt.Println(*repos[i].CloneURL)
	}





	// exec.Command("ebb", filename).Run()
	return 0
}

func main() {

	os.Exit(run());

}
