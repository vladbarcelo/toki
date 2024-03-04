# toki

## What is this project

This is just a simple CLI tool for filtering JSON logs that has a syntax similar to Loki.

## How to use

```bash
$ toki [logfile]
```

## Syntax

```
[col1,col2,col3] -> { key: "value" } |= "search query" ~last5m
 ^ the list of        ^ search by        ^ raw search   ^ time interval
   columns to           JSON values        query
   show                 in log
```

Example (filter nest.js error logs related to the database in the last 5 minutes):
```
[level,time,msg] -> { level: "error" } |= "database" ~last5m
```