package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"testing"
	"time"
)

func TestUploader(t *testing.T) {
	runtime.MemProfileRate = 1
	client, err := NewS3Client(os.Getenv("TEST_PROFILE"))
	if err != nil {
		t.Fatal(err)
	}
	const (
		srcBucket = "doollee-backup-bucket"
		srcKey    = "largeBackup.rdb"
		dstBucket = "adler-test"
		dstKey    = "bigbig.rdb"
	)

	ctx, cancel1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel1()
	ctx, cancel2 := context.WithTimeout(ctx, 30*time.Second)
	defer cancel2()
	err = CopyObjectClientSide(ctx, client, dstBucket, dstKey, client, srcBucket, srcKey)
	if err != nil {
		t.Fatal(err)
	}
}
