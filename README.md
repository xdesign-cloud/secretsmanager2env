# secretsmanager2env
This is a utility to take a secret from AWS Secrets Manager, and print a series of `export` statements to be loaded into a shell.

The envisioned use case for this is for containerising applications which aren't written in such a way that can be extended with Secrets Manager support.

## Usage
Run the application with `secretsmanager2env path/to/secret` to receive a series of `export` statements.

## Example
Assuming a secret with the following content:
```json
{
    "foo": "foobar",
    "bar_foo": "barfoo123"
}
```

This can be used in conjunction with the `source` command like so:
```
$ source <(secretsmanager2env example/secret)
$ echo $FOO
foobar

$ echo $BAR_FOO
barfoo123
```
