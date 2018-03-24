# go-prestamp

---

A simple CLI that generates `.go` files with custom *"header comments"*.


**Example usage:**

```bash
stjepan@carthaginian:~$ go-prestamp --author me --project sth    first second.go third

stjepan@carthaginian:~$ ls
first.go    second.go   third.go

stjepan@carthaginian:~$ cat first.go
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*
*           Author:  me
*           Project: sth
*           (c) 2018
*
*
*           Simple, but not simpler. - Albert Einstein
*
*
*           File created on Fri 23.3.2018 at 21:56
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
```
