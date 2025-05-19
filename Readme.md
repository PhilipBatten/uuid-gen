# go uuid gen library

I needed a quick method of getting a uuid4 in my clipboard to be able to setup mock data etc

This works for me but may not work for you.
Setup:
- popos
- xclip

This is untested for macos although some code for it was included.

How to use:
Add alias to .bashrc
```
alias uuid4="~/repos/uuid-gen/bin/uuid-gen"
```

Run uuid4 in terminal

Build:
```
go build -o bin
```