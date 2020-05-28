# Installation
* Download the binary appropriate to your architecture from https://github.com/figadore/automfa/releases
* make the file executable (e.g. `chmod +x automfa_darwin_amd64`)
* Move the file to a folder in your $PATH (e.g. `mv automfa_darwin_amd64 /usr/local/bin/automfa`)


# Usage

Add a key to the keyring with `automfa -a [-c] <service>` where `service` is replaced with what you want to nickname this generator (e.g. `vip` or `okta`)

Run `automfa <service>` to get a TOTP
