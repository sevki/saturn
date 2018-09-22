package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bcampbell/fuzzytime"
	"sevki.org/x/pretty"
)

var tf = `---
title: %s
author:
- name: Author One
  affiliation: University of Somewhere
- name: Author Two
  affiliation: University of Nowhere
date: Mon Apr  9 19:02:45 CES 2018


tags: [Performance, protobuf]
abstract: |
  This is a guide handy guide for improving your applications performance.
--- 

## What causes slow performance?
Slow performance can be caused by a variety of things from connection speeds
to data encoding.
 


asdasd`

func TestParseFile(t *testing.T) {
	title := "some Title"

	ft, _, err := fuzzytime.Extract(string("Mon Apr  9 19:02:45 CES 2018"))
	if err != nil {
		log.Println(err)
	}

	tym, _ := time.Parse(fuzzyFormat, ft.String())
	buf := bytes.NewBufferString(fmt.Sprintf(
		tf,
		title,
	))
	post := parseHeader(buf)
	if post.Title != title {
		t.Fail()
	}
	if !post.Date.Equal(tym) {
		t.Logf("was expecting %q got %q instead", tym, post.Date)
		t.Fail()
	}
	t.Log(pretty.JSON(post))
}