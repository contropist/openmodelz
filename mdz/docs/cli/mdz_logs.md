## mdz logs

Print the logs for a deployment

### Synopsis

Print the logs for a deployment

```
mdz logs [flags]
```

### Examples

```
  mdz logs blomdz-560m
```

### Options

```
  -e, --end string     Only return logs before this timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
  -h, --help           help for logs
  -s, --since string   Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes) (default "2006-01-02T15:04:05Z")
  -t, --tail int       Number of lines to show from the end of the logs
```

### Options inherited from parent commands

```
      --debug        Enable debug logging
  -u, --url string   URL to use for the server (MDZ_URL) (default http://localhost:80)
```

### SEE ALSO

* [mdz](mdz.md)	 - mdz manages your deployments

###### Auto generated by spf13/cobra on 26-Jul-2023