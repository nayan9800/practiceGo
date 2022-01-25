package gogit

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Commit struct {
	Hash       string
	Author     string
	AuthorName string
	When       string
	Message    string
}

var (
	gitLog = log.New(os.Stdout, "GITLOG:", log.LstdFlags|log.Lshortfile)
)

/*Get all Commits from head ordered in commit time*/
func GetAllComits(repo *git.Repository) ([]Commit, error) {
	commits := []Commit{}
	head, err := repo.Head()
	if err != nil {
		return commits, err
	}
	comitIter, err := repo.Log(&git.LogOptions{From: head.Hash(),
		All:   true,
		Order: git.LogOrderCommitterTime})

	comitIter.ForEach(func(c *object.Commit) error {
		commits = append(commits,
			Commit{
				Hash:       c.Hash.String(),
				Author:     c.Author.Email,
				AuthorName: c.Author.Name,
				When:       c.Committer.When.Format(time.RFC1123Z),
				Message:    c.Message,
			},
		)
		return nil
	})
	return commits, err
}

/*load the given repository in memory*/
func LoadRepo(path string) *git.Repository {
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: path,
	})
	if err != nil {
		gitLog.Println(err.Error())
	}
	return repo
}

//TODO: git plain clone
func Clone(path, url string) (*git.Repository, error) {
	name := getGitRepoName(url)
	repo, err := git.PlainClone(filepath.Join(path, name), false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}
	return repo, nil
}
func getGitRepoName(repo string) string {

	tokens := strings.Split(repo, "/")
	name := strings.Split(tokens[len(tokens)-1], ".")
	return name[0]
}

//TODO: git chekout branch
func ChekOutBranch(name string, repo *git.Repository) *git.Repository {

	var r plumbing.Reference
	var found bool
	s := repo.Storer
	refs, err := s.IterReferences()
	if err != nil {
		gitLog.Println(err.Error())
	}
	allBranches := getListAllRemotes(refs)
	for _, v := range allBranches {
		if strings.Contains(v.Name().String(), name) {
			found = true
			r = *v
			gitLog.Println(v.Name().Short(), v.Hash())
		}
	}
	if !found {
		return repo
	}
	w, err := repo.Worktree()
	if err != nil {
		gitLog.Println(err.Error())
		return repo
	}
	err = w.Checkout(&git.CheckoutOptions{Branch: r.Name(), Force: true})
	if err != nil {
		gitLog.Println(err.Error())
	}
	return repo
}
func getListAllRemotes(refs storer.ReferenceIter) []*plumbing.Reference {

	allRemotes := []*plumbing.Reference{}
	iter := storer.NewReferenceFilteredIter(isRemote, refs)
	iter.ForEach(func(r *plumbing.Reference) error {
		allRemotes = append(allRemotes, r)
		return nil
	})
	return allRemotes
}
func isRemote(ref *plumbing.Reference) bool {
	return ref.Name().IsRemote()
}

//TODO: git log
/*func Log(repo *git.Repository) {

	ref, err := repo.Head()
	if err != nil {
		gitLog.Println(err.Error())
	}
	citer, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Println(err.Error())
	}
	citer.ForEach(func(c *object.Commit) error {

		gitLog.Println(c.Hash, " ", c.Message)
		return nil
	})

}*/

//TODO: git cherry pick up
func CherryPickUp(repo *git.Repository) {
}
func TestGogit() {

	repo, err := Clone("./testdata", "https://github.com/fatih/color.git")
	if err == git.ErrRepositoryAlreadyExists {
		repo = LoadRepo("./testdata/color")
	}
	h, _ := repo.Head()
	gitLog.Println(h.Hash())
	repo = ChekOutBranch("update-ci", repo)
	h, _ = repo.Head()
	gitLog.Println(h.Hash())

}
