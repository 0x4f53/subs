# subs

Grab valid domains and subdomains from files, split them if they're fused and more - straight from the command-line! 
This is a command-line utility written in Golang as a proof-of-concept for the [textsubs package](https://github.com/0x4f53/textsubs).

### Installation
##### Linux and macOS

Simply run the `./install.sh` script (don't 
have the time to put this on package managers)

```bash
chmod +x install.sh
./install.sh
```

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

---

Copyright (c) 2024  Owais Shaikh

Licensed under [GNU GPL 3.0](LICENSE)