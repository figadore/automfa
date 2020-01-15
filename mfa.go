package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"

	"github.com/creachadair/otp"
	"github.com/zalando/go-keyring"
)

type options struct {
	account string
	service string
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		log.Fatalf("Error parsing args: %v", err)
	}
	cfg := otp.Config{
		Hash:   sha256.New,
		Digits: 6,
	}
	keyring_key := fmt.Sprintf("automfa_%s", opts.service)
	secret, err := keyring.Get(keyring_key, opts.account)
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}
	if err := cfg.ParseKey(secret); err != nil {
		log.Fatalf("Parsing key: %v", err)
	}
	fmt.Println(cfg.TOTP())
}

// Handle any complicated arg parsing here
func parseArgs() (*options, error) {
	// go-keyring module uses two keys to look up a secret, 'account' and
	// 'service'. This appears to have been done for cross-platform
	// compatibility, so it seems best to hardcode the account, and store the
	// username and password as two separate secrets
	opts := options{}
	opts.account = "automfa"
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return &options{}, fmt.Errorf("expecting at least 1 arg (secret_name), found %d", len(args))
	}
	opts.service = args[0]
	return &opts, nil
}
