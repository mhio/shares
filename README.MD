# shares

Generate base64 encoded shamir's secret sharing keys.

```
$ ./shares-linux-amd64 -h
Usage of ./shares-darwin-amd64:
  -k int
    	Number of keys to generate
  -l int
    	Length of generated key (default 64)
  -t int
    	Threshold # of keys to recreate
```

```
$ ./shares-linux-amd64 -k 5 -t 3 -l 32
random key[32]
kcYKTr52jK60FQx5/YKAyzlLqrZ9tWHHZM84cIXIfZA=

generating keys[5] threshold[3]
key[0]: fq6IaeCN+EvF816i0SDzdUcwpX3iIZ+s5n0Yw+22+dAp
key[1]: lBCdbwHKbsUjrQh4zXDFvOIc0m9whugxDR8aRjZMB4Tm
key[2]: gW0YJSLRUWp8gsH1XPrhCWH1pkzvkeuM4tsNrSrG/8wx
key[3]: n216h6zKmVKxb9Zetw2+8pUqmKzZiCaLFWPwY2MRjYys
key[4]: GkOtKqW2m2oA6MEXToqoLxlAxuSJhjgMitRsXNVcEYHe

Using [3] keys ok: kcYKTr52jK60FQx5/YKAyzlLqrZ9tWHHZM84cIXIfZA=
Using [4] keys ok: kcYKTr52jK60FQx5/YKAyzlLqrZ9tWHHZM84cIXIfZA=
Using [5] keys ok: kcYKTr52jK60FQx5/YKAyzlLqrZ9tWHHZM84cIXIfZA=
```
