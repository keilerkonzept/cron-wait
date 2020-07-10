# ${APP}

A tiny tool that waits until a given cron expression would trigger, and then just exits. If multiple expressions are given, it waits until the _first_ (earliest) match.

## Example


```sh
$ ${APP} "*/5 * * * * *"
2020/07/08 17:57:17 waiting 2.998536s until next match (2020-07-08T17:57:20+02:00) of cron expression "*/5 * * * * *"
$ # 2.998s later
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
