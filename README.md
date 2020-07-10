# wait-for-cron-expression-match

A tiny tool that waits until a given cron expression would trigger, and then just exits. If multiple expressions are given, it waits until the _first_ (earliest) match.

## Example


```sh
$ wait-for-cron-expression-match "*/5 * * * * *"
2020/07/08 17:57:17 waiting 2.998536s until next match (2020-07-08T17:57:20+02:00) of cron expression "*/5 * * * * *"
$ # 2.998s later
```

## Contents

- [Get it](#get-it)
- [Usage](#usage)

## Get it

Using go get:

```bash
go get -u github.com/keilerkonzept/wait-for-cron-expression-match
```

Or [download the binary for your platform](https://github.com/keilerkonzept/wait-for-cron-expression-match/releases/latest) from the releases page.

## Usage

```text
wait-for-cron-expression-match [OPTIONS] [CRON_EXPRESSION [CRON_EXPRESSIONS...]]

Usage of wait-for-cron-expression-match:
  -dots
    	Print dots to stdout while waiting
  -q	(alias for -quiet)
  -quiet
    	Suppress all output
```
