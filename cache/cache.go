package cache

import (
	colly "colly/crawler"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var Rows *DB = &DB{Indexer: Indexer{
	StringIndex: make(map[string]*[]*Row),
	RegexpIndex: make(map[string]*[]*Row),
}}

type Row struct {
	Title, Link string
}
type Indexer struct {
	StringIndex map[string]*[]*Row
	RegexpIndex map[string]*[]*Row
}

type DB struct {
	mu      sync.Mutex
	Storage map[string]string
	Indexer
}

func init() {
	if err := deserialization(); err != nil {
		Rows.Storage = colly.Scrape()
		b, _ := json.Marshal(Rows.Storage)
		os.WriteFile("./cache/dump.json", b, os.ModePerm)
	}
}

func deserialization() error {
	raw, err := os.ReadFile("./cache/dump.json")
	if err != nil {
		log.Printf("deserialization loadfailure :%s: \n", err)
		return err
	}
	err = json.Unmarshal(raw, &Rows.Storage)
	if err != nil {
		log.Printf("deserialization unmarshalfailure :%s: \n", err)
		return err
	}
	return nil
}

func (idx *Indexer) Hit(t string) (hits []*Row, missed bool) {
	stg := idx.StringIndex
	if r, exist := stg[t]; exist {
		hits = *r
		return
	} else {
		for ft, h := range stg {
			// `abc` use `a`
			if strings.Contains(t, ft) {
				for _, r := range *h {
					if strings.Contains(r.Title, t) {
						hits = append(hits, r)
					}
				}
				stg[t] = &hits
				return
			}
		}
		missed = true
		return
	}
}
func (idx *Indexer) RegexpHit(t string) (hits []*Row, missed bool) {
	stg := idx.StringIndex
	if r, exist := stg[t]; exist {
		hits = *r
		return
	}
	missed = true
	return
}

func (db *DB) Search(t string) []*Row {
	db.mu.Lock()
	defer db.mu.Unlock()
	if hs, m := db.Hit(t); m {
		var hits []*Row = nil
		for tl, l := range db.Storage {
			if strings.Contains(tl, t) {
				hits = append(hits, &Row{tl, l})
			}
		}
		db.Indexer.StringIndex[t] = &hits
		return hits
	} else {
		fmt.Println("using index")
		return hs
	}
}

func (db *DB) SearchWithRegexp(reg string) ([]*Row, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if hs, m := db.Hit(reg); m {
		var hits []*Row = nil
		for tl, l := range db.Storage {
			if matched, err := regexp.MatchString(reg, tl); err != nil {
				return nil, err
			} else if matched {
				hits = append(hits, &Row{tl, l})
			}
		}
		db.Indexer.StringIndex[reg] = &hits
		return hits, nil
	} else {
		fmt.Println("using index")
		return hs, nil
	}

}
