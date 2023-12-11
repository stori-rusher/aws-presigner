# aws-presigner

A simple command-line tool to create AWS S3 presigned URLs. The main motivation to create this was
because creating PUT URLs is not possible via the console or AWS CLI tool.

## Usage

To use, download the executable file and run like this:

```sh
./presign -b rusher-test -k some-file.txt -m put
```

If you get a permission error, you made need to update the file permissions

```sh
chmod +x presign
```

## Switching environments

This tool loads the default AWS config, so it relies on whichever AWS CLI profile is currently set in your enviroment. To use a different environment without globally changing your aws profile, you can prepend the command to set the `AWS_PROFILE` environment variable for the execution

```sh
AWS_PROFILE=qa ./presign -b rusher-test -k some-file.txt -m put
```
