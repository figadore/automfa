# Installation
* Download the binary appropriate to your architecture from https://github.com/figadore/automfa/releases
* make the file executable (e.g. `chmod +x automfa_darwin_amd64`)
* Move the file to a folder in your $PATH (e.g. `mv automfa_darwin_amd64 /usr/local/bin/automfa`)

## Mac additional security
You may have to run the following to get around the error about the developer not being verified
```
sudo spctl --master-disable
automfa -h
sudo spctl --master-enable
```


# Usage
See 
```
automfa --help
```

Add a key to the keyring with `automfa -a [-c] <service>` where `service` is replaced with what you want to nickname this generator (e.g. `vip` or `okta`). `-c`

Run `automfa <service>` to get a TOTP
