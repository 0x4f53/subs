# subs

Grab valid domains and subdomains from files, split them if they're fused and more - straight from the command-line! 
This is a command-line utility written in Golang as a proof-of-concept for the [textsubs package](https://github.com/0x4f53/textsubs).

Note: This tool only extracts subs from text. It does NOT resolve them. That is a time-consuming process which needs
to be done manually with several ethical considerations.

### Installation
##### Linux and macOS

Simply run the `./install.sh` script (don't 
have the time to put this on package managers)

```bash
chmod +x install.sh
./install.sh
```

You can also find the binaries in [`.build/binaries`](.build/binaries/) if you want to directly run them
without installation

### Usage
```bash
subs [input_file] [flags]
```

### Flags
  -d, --domains   Get domains only
  -h, --help      Help
  -u, --unique    Only get unique entries

### Examples

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
❯ wget "https://crt.sh/?q=129341" -O .temp && subs .temp -u -p && rm .temp
{"subdomain":"crt.sh","domain":"crt.sh"}
{"subdomain":"fonts.googleapis.com","domain":"googleapis.com"}
{"subdomain":"ct.googleapis.com","domain":"googleapis.com"}
{"subdomain":"plausible.ct.nordu.net","domain":"nordu.net"}
...
```

---

Copyright (c) 2024  Owais Shaikh

Licensed under [GNU GPL 3.0](LICENSE)