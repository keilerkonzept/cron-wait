# ${APP}

A tiny tool that waits until a given cron expression would trigger, and then just exits. If multiple expressions are given, it waits until the _first_ (earliest) match.

## Example


```sh
$ ${APP} "*/5 * * * *"
[wait-for-cron-expression-match] 2020/07/10 11:19:46.883743 waiting 13.116478s until next match (2020-07-10T11:20:00+02:00) of cron expression ["*/5 * * * *"]
$ # 13s later
```

## Contents

- [Get it](#get-it)
- [Usage](#usage)

## Get it

Using go get:

```bash
go get -u github.com/keilerkonzept/${APP}
```

Or [download the binary for your platform](https://github.com/keilerkonzept/${APP}/releases/latest) from the releases page.

## Usage

```text
${APP} [OPTIONS] [CRON_EXPRESSION [CRON_EXPRESSIONS...]]

${USAGE}
```
