package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	gogit "github.com/nayan9800/practiceGo/pkg/goGit"
	"github.com/spf13/cobra"
)

//loging template string
var logStr = "{{range .}}\u001b[33mcommit {{.Hash}}\u001b[0m\nAuthor: {{.AuthorName}} <{{.Author}}>\nDate: {{.When}}\n\n{{.Message}}\n{{end}}"
var (
	//cmd logger
	cmdLog = log.New(os.Stdout, "CMDLOG: ", log.LstdFlags|log.Lshortfile)
	/*commands*/
	rootcmd = &cobra.Command{
		Use:   "mygit",
		Short: "mygit is copy of Git CVS",
	}

	cloneCmd = &cobra.Command{
		Use:   "clone <repo>",
		Short: "clone git repository in given path",
		RunE:  cloneRun,
	}

	logCmd = &cobra.Command{
		Use:   "log",
		Short: "log's the commit history",
		RunE:  logRun,
	}

	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "push changes to remote",
		Run:   pushRun,
	}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize a new or convert an existing, unversioned project to a Git repository",
		Run:   initRun,
	}
)

/*initialize a new or convert an existing,
unversioned project to a Git repository*/
func initRun(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("path")
	isbare, _ := cmd.Flags().GetBool("bare")
	_, err := git.PlainInit(path, isbare)
	if err != nil {
		log.Println(err.Error())
	}
}

/*pushes code to remote repository*/
func pushRun(cmd *cobra.Command, args []string) {
	repo := gogit.LoadRepo(".")
	if err := gogit.Gitpush(repo); err != nil {
		log.Println(err.Error())
	}
}

/*log the commit history of git repository*/
func logRun(cmd *cobra.Command, args []string) (err error) {

	path, _ := cmd.Flags().GetString("path")
	isoneline, _ := cmd.Flags().GetBool("oneline")
	repo := gogit.LoadRepo(path)

	if !isoneline {
		c, err := gogit.GetAllComits(repo)
		if err != nil {
			return err
		}
		lt, _ := template.New("log").Parse(logStr)
		err = lt.Execute(os.Stdout, c)
		if err != nil {
			return err
		}
	} else {
		head, err := repo.Head()
		comitIter, err := repo.Log(&git.LogOptions{From: head.Hash(),
			All:   true,
			Order: git.LogOrderCommitterTime})

		comitIter.ForEach(func(c *object.Commit) error {

			ans := strings.Join(strings.Fields(c.Message), " ")
			fmt.Printf("\u001b[33m%s\u001b[0m %s\n", c.Hash.String()[:7], ans)
			return err
		})
	}
	return
}

/*clone git repository in given path*/
func cloneRun(cmd *cobra.Command, args []string) (err error) {
	path, _ := cmd.Flags().GetString("path")
	if len(args) == 0 || len(args) > 1 {
		cmd.Help()
		err = errors.New("not enough arguments")
	}
	if _, err := gogit.Clone(path, args[0]); err != nil {
		return err
	}
	return
}
func Execute() {
	//disable default cmd
	rootcmd.CompletionOptions.DisableDefaultCmd = true
	//add flags to inicmd
	initCmd.Flags().Bool("bare", false, "Create a bare repository")
	initCmd.Flags().StringP("path", "p", ".", "path of working directory")
	//add flags to logcmd
	logCmd.Flags().StringP("path", "p", ".", "path of working directory")
	logCmd.Flags().Bool("oneline", false, "log in oneline")
	//add flags to clone cmd
	cloneCmd.Flags().StringP("path", "p", ".", "path of working directory")

	//add all commands to rootcmd
	rootcmd.AddCommand(cloneCmd)
	rootcmd.AddCommand(logCmd)
	rootcmd.AddCommand(pushCmd)
	rootcmd.AddCommand(initCmd)

	//execute rootcmd
	if err := rootcmd.Execute(); err != nil {
		cmdLog.Fatal(err.Error())
	}
}
