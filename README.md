# launeparser
Parsing newspapers everyday to get a corpus

This program periodically scrapes web pages, two times a day and dumps the
resulting texts to a directory with this format :

```
out/
   newspaper name/
      2018-04-09_21:00.txt
```

You can specify as many sites to scrapes as you want in the configuration file.

## Usage

```
Launeparser scrapes newspapers

Usage:
  launeparser [command]

Available Commands:
  help        Help about any command
  scrape      Instantly scrape
  start       Start the server and scraping
  version     Show build and version

Flags:
  -h, --help                help for launeparser
      --log.format string   one of text or json (default "text")
      --log.level string    one of debug, info, warn, error or fatal (default "info")
      --log.line            enable filename and line in logs
      --output string       output directory (default "out")

Use "launeparser [command] --help" for more information about a command.
```

## Configure

```yaml
server:
  host: 127.0.0.1
  port: 8012
  debug: true
log:
  level: debug
  format: text
  line: true
newspapers:
  - url: http://...
    name: ...
```

The `server` part is not needed, as well as the `log` server as there are sane
defaults. Also the `server` part is completely unused when using the 
`launeparser scrape` command.

