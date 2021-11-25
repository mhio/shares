# shares

Generate base64 encoded shamir's secret sharing keys.

```
$ ./shares-linux-amd64 -h
Usage of ./shares-linux-amd64:
  -k int
      # of keys
  -l int
      Length of generated key (default 64)
  -t int
      Threshold # of keys to recreate
```

```
$ ./shares-linux-amd64 -k 5 -t 3 -l 32
random key[32]
ip9rb+kGqrLkXwlrzFSoXOVY+OW7rn77eaXKfEDh7tc=
generating keys[5] threshold[3]
key[0]: OF73jXGem3MIaQqkcKw+sbgpNqHy1DIx+SqvDpPpTcfe
key[1]: sBzEMAPUg8BrQMueWw5xVGme/Y/ZgpKmkgI4Xgi28FxR
key[2]: t83PcLwEJ8wkBo0LTbX7w9mmzetv+scaB1nR6qKx4QVT
key[3]: m37Oz5nKrwn3V7oPzzpAhVbbr+UkblZ81LWpGJgGIrvb
key[4]: LO0KkNJdUqUJ/GIodmjMrmL1aPtbJaG6QayWh5DPxWDQ
Using [3] keys [2 3 0]
ip9rb+kGqrLkXwlrzFSoXOVY+OW7rn77eaXKfEDh7tc=
Using [4] keys [4 0 2 1]
ip9rb+kGqrLkXwlrzFSoXOVY+OW7rn77eaXKfEDh7tc=
Using [5] keys [4 1 3 0 2]
ip9rb+kGqrLkXwlrzFSoXOVY+OW7rn77eaXKfEDh7tc=
```
