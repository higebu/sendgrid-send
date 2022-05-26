# sendgrid-send

A Simple SendGrid Test Tool

# Install

```shell
go install github.com/higebu/sendgrid-send
```

# Usage

```shell
export SENDGRID_API_KEY=xxxxx
sendgrid-send -to test@example.com -to-name 'Example User' \
	-from test@example.com -from-name 'Example User' \
	-content 'Hello SendGrid'
```
