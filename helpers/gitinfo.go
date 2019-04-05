package helpers

import (
	"github.com/tsuyoshiwada/go-gitlog"
	"time"
)

type GitPageInfo struct {
	AbbreviatedHash string                       `json:"abbreviated_hash"`
	AuthorName      string                       `json:"author_name"`
	AuthorEmail     string                       `json:"author_email"`
	AuthorDate      time.Time                    `json:"author_date"`
	Hash            string                       `json:"hash"`
	Subject         string                       `json:"subject"`
	Commits         []*gitlog.Commit             `json:"commits"`
	Contributors    map[string]*gitlog.Committer `json:"contributors"`
}
