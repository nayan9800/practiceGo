package cmd

import (
	"errors"
	"log"
	"os"
	"text/template"

	gogit "github.com/nayan9800/practiceGo/pkg/goGit"
	"github.com/spf13/cobra"
)

var logStr = "{{range .}}\u001b[33mcommit {{.Hash}}\u001b[0m\nAuthor: {{.AuthorName}} <{{.Author}}>\nDate: {{.When}}\n\n{{.Message}}\n{{end}}"
var (
	cmdLog  = log.New(os.Stdout, "CMDLOG: ", log.LstdFlags|log.Lshortfile)
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
)

func pushRun(cmd *cobra.Command, args []string) {
	repo := gogit.LoadRepo(".")
	if err := gogit.Gitpush(repo); err != nil {
		log.Println(err.Error())
	}
}
func logRun(cmd *cobra.Command, args []string) (err error) {

	path, _ := cmd.Flags().GetString("path")
	repo := gogit.LoadRepo(path)
	c, err := gogit.GetAllComits(repo)
	lt, err := template.New("log").Parse(logStr)
	err = lt.Execute(os.Stdout, c)
	return
}
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
	logCmd.Flags().StringP("path", "p", ".", "path of working directory")
	cloneCmd.Flags().StringP("path", "p", ".", "path of working directory")
	rootcmd.AddCommand(cloneCmd)
	rootcmd.AddCommand(logCmd)
	rootcmd.AddCommand(pushCmd)
	if err := rootcmd.Execute(); err != nil {
		cmdLog.Fatal(err.Error())
	}
}
