You Can only control one key to trace all key used in normal internet.

```shell
go clone https://github.com/r1cs/passwordGenerator
go build -o $GOBIN/passwordGenerator ./cmd/
passwordGenerator --url github.com --name test --key haha 
```

#### note
shell will remain history, which causing the leak of your origin key.So use `git bash shell` in windows instead, and clean the history
```shell
cat /dev/null > ~/.bash_history
chmod -w ~/.bash_history
```
now git bash history will not recorded anymore
