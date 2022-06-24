# GoBackup is a fullstack backup tool 
design for web servers similar with [backup/backup](https://github.com/backup/backup), work with Crontab to backup automatically.

## Features

- No dependencies.
- Multiple Databases source support.
- Multiple Storage type support.
- Archive paths or files into a tar.

## Current Support status

### Databases

- MySQL
- PostgreSQL
- Redis - `mode: sync/copy`
- MongoDB

### Archive

Use `tar` command to archive many file or path into a `.tar` file.

### Compressor

- Tgz - `.tar.gz`
- Uncompressed - `.tar`

### Encryptor

- OpenSSL - `aes-256-cbc` encrypt

### Storages

- Local
- FTP
- SCP - Upload via SSH copy
- [Amazon S3](https://aws.amazon.com/s3)

## Install (macOS / Linux)

```bash
$ curl -sSL https://git.io/gobackup | bash
```

after that, you will get `/usr/local/bin/gobackup` command.

```bash
$ gobackup -h
NAME:
   gobackup - Easy full stack backup operations on UNIX-like systems

USAGE:
   gobackup [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     perform
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Configuration

GoBackup will seek config files in:

- ~/.gobackup/gobackup.yml
- /etc/gobackup/gobackup.yml

Example config: [gobackup_test.yml](https://github.com/4nkitd/gobackup/blob/master/gobackup_test.yml)

```yml
# gobackup config example
# -----------------------
models:
  JobName:
    compress_with:
      type: tgz
    store_with:
      type: s3
      keep: 20
      bucket: gobackup-test
      region: ap-south-1
      path: backups
      access_key_id: Ohsgwk86h2ksas
      secret_access_key: Ojsiw729wujhKdhwsIIOw9173
    databases:
      DbName:
        type: mysql
        host: localhost
        port: 3306
        database: test
        username: root
        password: 123456
```

## Backup Run

You may want run backup in scheduly, you need Crontab:

```bash
gobackup perform >> ~/.gobackup/gobackup.log
```
And after a day, you can check up the execute status by `~/.gobackup/gobackup.log`.

