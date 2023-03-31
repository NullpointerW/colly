package cache

import (
	colly "colly/crawler"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

var Rows *DB=&DB{}

type Row struct {
	Title, Link string
}
type Indexer struct {
	Index map[string]*[]Row
}

type DB struct {
	mu      sync.Mutex
	storage map[string]string
	Indexer
}

func init() {
	if err := deserialization(); err != nil {
		Rows.storage = colly.Scrape()
		b, _ := json.Marshal(Rows.storage)
		os.WriteFile("./cache/dump.json", b, os.ModePerm)
	}
}

func deserialization() error {
	raw, err := os.ReadFile("./cache/dump.json")
	if err != nil {
		log.Printf("deserialization loadfailure :%s: \n", err)
		return err
	}
	err = json.Unmarshal(raw, &Rows.storage)
	if err != nil {
		log.Printf("deserialization unmarshalfailure :%s: \n", err)
		return err
	}
	return nil
}

func (idx *Indexer) Hit(t string) (hits []Row, missed bool) {
	stg := idx.Index
	if r, exist := stg[t]; exist {
		hits = *r
		return
	} else {
		for ft, h := range stg {
			if strings.Contains(ft, t) {
				stg[t] = h
				hits = *h
				return
			}
		}
		missed = true
		return
	}
}

func (db *DB) Search(t string) []Row {
	db.mu.Lock()
	defer db.mu.Unlock()
	if hs, m := db.Hit(t); m {
		var hits []Row = nil
		for tl, l := range db.storage {
			if strings.Contains(tl, t) {
				hits = append(hits, Row{tl, l})
			}
			if len(hits) != 0 {
				db.Indexer.Index[t] = &hits
			}
			return hits
		}
	} else {
		return hs
	}
	return nil
}
