package main

import (
	"kkblogs/blogs-api/common"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	password := common.GetPassword("admin-kk", "uyTZFr")
	log.Println(password)
}
