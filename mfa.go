package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"

	"github.com/creachadair/otp"
	"github.com/zalando/go-keyring"
)

type options struct {
	add       bool
	account   string
	service   string
	clearText bool
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		log.Fatalf("Error parsing args: %v", err)
	}
	if opts.add {
		addService(opts)
		return
	}
	cfg := otp.Config{
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

func addService(opts *options) {
	var secret string
	var e error
	fmt.Printf("Enter shared-secret (aka MFA seed) for %s:", opts.service)
	if opts.clearText {
		_, e = fmt.Scanln(&secret)
	} else {
		var secretBytes []byte
		secretBytes, e = terminal.ReadPassword(int(os.Stdin.Fd()))
		secret = string(secretBytes)
		// Start new line after prompt
		fmt.Println("")
	}

	if e != nil {
		log.Fatalln("Error reading secret from stdin:", e)
	}
	keyring_key := fmt.Sprintf("automfa_%s", opts.service)
	e = keyring.Set(keyring_key, opts.account, secret)
	if e != nil {
		// Secret is in memory, but unable to write it to the OS keyring. Continue with warning
		log.Println("Error writing secret:", e)
	}
}

// Handle any complicated arg parsing here
func parseArgs() (*options, error) {
	// go-keyring module uses two keys to look up a secret, 'account' and
	// 'service'. This appears to have been done for cross-platform
	// compatibility, so it seems best to hardcode the account, and store the
	// username and password as two separate secrets
	opts := options{}
	opts.account = "automfa"
	flag.BoolVar(&opts.add, "a", false, "Add a new MFA secret")
	flag.BoolVar(&opts.clearText, "c", false, "If set, show text while typing (useful where copy/paste unavailable)")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return &options{}, fmt.Errorf("expecting at least 1 arg (service), found %d", len(args))
	}
	opts.service = args[0]
	return &opts, nil
}
