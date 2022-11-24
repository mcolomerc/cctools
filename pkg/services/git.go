package services

import (
	"mcolomerc/cc-tools/pkg/config"
	"mcolomerc/cc-tools/pkg/log"

	"github.com/go-git/go-git/v5"
)

type GitService struct {
	Repos  map[string]string
	Output string
}

func NewGitService(config config.Config) *GitService {
	return &GitService{
		Repos:  config.Export.Git,
		Output: config.Export.Output,
	}
}

func (g *GitService) Export() {
	log.Debug("Cloning repositories...")
	done := make(chan bool, len(g.Repos))
	for name, addr := range g.Repos {
		go func(url string, output string) {
			g.Clone(url, output)
			done <- true
		}(addr, g.Output+"/"+name)
	}
	for i := 0; i < len(g.Repos); i++ {
		<-done
	}
}

func (g *GitService) Clone(url string, output string) error {
	log.Debug("Cloning ... %s", url)
	_, err := git.PlainClone(output, false, &git.CloneOptions{URL: url})
	if err != nil {
		log.Error("Cloning %s error", url)
		log.Error(": ", err)
		return err
	}
	return nil
}
