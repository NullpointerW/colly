package cache

import (
	colly "colly/crawler"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var Rows *DB = &DB{Indexer: Indexer{Index: make(map[string]*[]*Row)}}

type Row struct {
	Title, Link string
}
type Indexer struct {
	Index map[string]*[]*Row
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
	stg := idx.Index
	if r, exist := stg[t]; exist {
		hits = *r
		return
	} else {
		for ft, h := range stg {
			// `abc` use `a`
			if strings.Contains(t, ft) {
				for   _,r:=  range *h{
                    if  strings.Contains(r.Title,t){
						hits=append(hits, r)
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
		
			db.Indexer.Index[t] = &hits
		
		return hits
	} else {
		fmt.Println("using index")
		return hs
	}
}
