package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	data := []struct {
		buf  []byte
		want int
	}{
		{
			buf:  []byte("mul(4,5)"),
			want: 20,
		},
		{
			buf:  []byte("xxxmul(4,5)"),
			want: 20,
		},
		{
			buf:  []byte("don't()mul(4,5)"),
			want: 0,
		},
		{
			buf:  []byte("don't()mul(4,5)do()mul(2,3)"),
			want: 6,
		},
		{
			buf:  []byte("don't()do()mul(4,5)do()mul(2,3)"),
			want: 26,
		},
		{
			buf:  []byte("do()mul(4,5)"),
			want: 20,
		},
		{
			buf:  []byte("don't()do()mul(4,5)"),
			want: 20,
		},
		{
			buf:  []byte("don't()do()mul(4,5)don't()"),
			want: 20,
		},
		{
			buf:  []byte("don't()do()mul(4,5)don't()do()"),
			want: 20,
		},
		{
			buf:  []byte("dddddon't()mul(4,5)don't()do()mul(5,5)"),
			want: 25,
		},
		{
			buf:  []byte("mul(3,4)dddddon't()mul(4,5)don't()mul(5,5)"),
			want: 12,
		},
	}

	for _, d := range data {
		sut := NewCursor(d.buf)
		got, err := sut.getTotalPt2()
		if err != nil {
			panic(err)
		}
		if got != d.want {
			log.Fatalf("TEST FAILED\n\twant: %d\n\tgot: %d", d.want, got)
		}

	}
}
