# wc
Unix wc tool

This is a Go implementation of the Unix wc tool. Inspired by [these challenges](https://codingchallenges.fyi/challenges/challenge-wc/).

How to use:
```
go build ccwc.go // compile
./ccwc -clw someFile.txt // count number bytes, lines, and words in the file
```

Can also read from stdin:
```
cat someFile.txt | ./ccwc -clw
```
