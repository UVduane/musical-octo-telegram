package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob" // driver for file://
	_ "gocloud.dev/blob/s3blob"   // driver for s3://
	//"gocloud.dev/gcerrors"
)

func main() {

	blobmux := blob.DefaultURLMux()
	ctx := context.TODO()
	u := os.Args[1]
	fn := os.Args[2]

	bucket, err := blobmux.OpenBucket(ctx, u)
	if err != nil {
		fmt.Printf("unable to open bucket %s: %v\n", u, err)
		os.Exit(1)
	}
	defer bucket.Close()

	fmt.Println("Bucket opened ok")

	reader, err := bucket.NewReader(ctx, fn, nil)
	if err != nil {
		fmt.Printf("unable to open file %s: %v\n", fn, err)
		os.Exit(1)
	}
	defer reader.Close()

	b, err := io.ReadAll(reader)
	if err != nil {
		fmt.Printf("unable to read from file %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)

}
