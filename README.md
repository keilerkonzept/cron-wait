# wait-for-cron-expression-match

A tiny tool that waits until a given cron expression would trigger, and then just exits. If multiple expressions are given, it waits until the _first_ (earliest) match.

## Example


```sh
$ wait-for-cron-expression-match "*/1 * * * *"
[wait-for-cron-expression-match] 2020/07/09 23:31:11.685480 waiting 48.314746s until next match (2020-07-09T23:32:00+02:00) of cron expression ["*/1 * * * *"]
# 48.3s later
$
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
wait-for-cron-expression-match [CRON_EXPRESSION [CRON_EXPRESSIONS...]]
```
