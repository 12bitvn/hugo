// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/tsuyoshiwada/go-gitlog"
	"path/filepath"
	"strings"
)

// RevNumber alias for `-n <number>`
type RevFile struct {
	FileName string
}

// Args ...
func (rev *RevFile) Args() []string {
	return []string{"-p", rev.FileName}
}

type gitInfo struct {
	contentDir string
	repo       gitlog.GitLog
}

func (g *gitInfo) forPage(p page.Page) *helpers.GitPageInfo {
	if g == nil {
		return nil
	}

	name := strings.TrimPrefix(filepath.ToSlash(p.File().Filename()), g.contentDir)
	name = strings.TrimPrefix(name, "/")
	logs, err := g.repo.Log(&RevFile{FileName: name}, &gitlog.Params{IgnoreMerges: true})
	if err != nil {
		return nil
	}
	if len(logs) == 0 {
		return nil
	}
	contributors := map[string]*gitlog.Committer{}
	for _, log := range logs {
		contributors[log.Author.Name] = log.Committer
	}
	return &helpers.GitPageInfo{
		AbbreviatedHash: logs[0].Hash.Short,
		AuthorName:      logs[0].Author.Name,
		AuthorEmail:     logs[0].Author.Email,
		AuthorDate:      logs[0].Author.Date,
		Hash:            logs[0].Hash.Long,
		Subject:         logs[0].Subject,
		Commits:         logs,
		Contributors:    contributors,
	}
}

func newGitInfo(cfg config.Provider) (*gitInfo, error) {
	workingDir := cfg.GetString("workingDir")
	git := gitlog.New(&gitlog.Config{
		Path: workingDir,
	})
	return &gitInfo{contentDir: workingDir, repo: git}, nil
}
