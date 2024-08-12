[![Golang](https://img.shields.io/badge/Golang-fff.svg?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-GNU%20GPL%203.0-purple?style=flat-square&logo=libreoffice)](LICENSE)
[![Latest Version](https://img.shields.io/github/v/tag/0x4f53/subs?label=Version&style=flat-square&logo=semver)](https://github.com/0x4f53/subs/releases)
[![Binaries](https://img.shields.io/badge/Binaries-Click%20Here-blue?style=flat-square&logo=dropbox)](.build/binaries/)

# 🔍 subs

<img src = preview.png alt="subs preview" width = "500dp">

A utility to grab valid domains and subdomains from files, split them if they're fused, resolve them and more - straight from the command-line! 
This is a command-line utility written in Golang as a proof-of-concept for the [textsubs package](https://github.com/0x4f53/textsubs).

Note: This tool only extracts subs from text. It does not extract URLs (there are several methods to do that, such as grep and regular
expressions)

Features:
- Splits fused strings (google.comapple.comblog.0x4f.in magically becomes google.com apple.com blog.0x4f.in)
- Resolves subdomains and domains concurrently in seconds
- Multiple kinds of output, including domain, subdomain or both as JSON!


## ⚙️ Building

To build this on your machine, you need to have Golang installed.
If you do, simply make build.sh executable and run it like so

```bash
❯ chmod +x build.sh
❯ ./build.sh
```

### ⚙️ Installation
##### Linux and macOS

Simply run the `./install.sh` script (don't 
have the time to put this on package managers)

```bash
chmod +x install.sh
sudo ./install.sh
```

And to uninstall

```bash
chmod +x uninstall.sh
sudo ./uninstall.sh
```

You can also find the binaries in [`.build/binaries`](.build/binaries/) if you want to directly run them
without installation

### 🖊️ Usage
```bash
subs [input_file] [flags]
```

#### Flags

  - `-h`, `--help`      Help
 
  - `-d`, `--domains`   Get domains only

  - `-r`, `--resolve`   Only get items that resolve (using local DNS settings)

  - `-p`, `--pair`   Pair as a JSON string {"subdomain":"www.example.com", "domain": "example.com"}
  
  - `-k`, `--keepDomains`   Keep both domains and subdomains in output (this catches both example.com and www.example.com)

  - `-r`, `--resolve`   Only get items that resolve (using local DNS settings)

#### Examples

Read a file on disk

```bash
❯ subs test.txt
subdomain1.example.com
subdomain2.example.com
subdomain3.example.com
subdomain4.example.com
...
```

Read all files in a directory

```bash
❯ for file in *; subs "$file"
www.gnu.org
google.golang.org
subdomain1.example.com
...
```

Get all the subdomains from a webpage and remove duplicates

```bash
❯ wget "https://crt.sh/?q=129341" -O .temp && subs .temp -p && rm .temp
{"subdomain":"crt.sh","domain":"crt.sh"}
{"subdomain":"fonts.googleapis.com","domain":"googleapis.com"}
{"subdomain":"ct.googleapis.com","domain":"googleapis.com"}
{"subdomain":"plausible.ct.nordu.net","domain":"nordu.net"}
...
```

Continuously scan certificates on crt.sh for alive subdomains with autoincremented id

```bash
❯ id=129341; while true; do wget "https://crt.sh/?q=$id" -O .temp && subs .temp -r -p >> output.txt && rm .temp; id=$((id + 1)); done
{"subdomain":"crt.sh","domain":"crt.sh"}
{"subdomain":"fonts.googleapis.com","domain":"googleapis.com"}
{"subdomain":"ct.googleapis.com","domain":"googleapis.com"}
{"subdomain":"plausible.ct.nordu.net","domain":"nordu.net"}
```

---

Copyright (c) 2024  Owais Shaikh

Licensed under [GNU GPL 3.0](LICENSE)