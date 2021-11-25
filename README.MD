# shares

Generate base64 encoded shamir's secret sharing keys.

```
$ ./shares-linux-amd64 -h
Usage of ./shares-darwin-amd64:
  -k int
    	Number of keys go generate
  -l int
    	Length of generated key (default 64)
  -t int
    	Threshold # of keys to recreate
```

```
$ ./shares-linux-amd64 -k 5 -t 3 -l 32
random key[64]
0Jlecf6ctgjx3RYujDsBon/zoJhcMU3u2aiuiVCaakS5axaAbXgbxRH7cw3WC7IUEbQlNLD+oRHrNYLiwwHg8A==

generating keys[6] threshold[4]
key[0]: beCTsV4i7S6MfURcPE54n2hq44lxjnYVPfkmvk7cAYA+aP4O5HzNqIcgB30XDlDirP+XqLtozJoBxS7DfPSzX/k=
key[1]: aawKIuKnceU8c5biKbRgXPlURXu2yJ7G5CYLJzbeNgzUDhMCeFN3CqPRpYinTs0Wo99GAX1QbrhO3VBsTiXxaKA=
key[2]: 44CbTxn5f6MUXcdbVMILJ6h1VwNvCyNfNqFfccdTPnQUDfnRBS6/az6gRglBGnsEr7aVN483t3rm1K2qbgyVQRI=
key[3]: LhcfPhXEqbf/genAsjWilzyQMPt6yLKTgKHblj9s/NzMw3e0dAzprAQcM8nr+M4aSgKYLBWO3LoJ7RGCgm5L8Js=
key[4]: G1QvE6ryaECQCVnjU6xp8WHaxoWhVBzDEXPi3TQ5YD70Jq4s1UdUJ5Y0h6e6hl/aG6Q+644M7t5U4heilUAMb2I=
key[5]: hToD8XOc8lJkMwq9cNpqMK0ddQm61y8cXIb6Q60LvoxM3X8eQmc+1eTc3aDoY33lf4NAgezGvKlLV6qYXJChk+s=

Using [4] keys ok: 0Jlecf6ctgjx3RYujDsBon/zoJhcMU3u2aiuiVCaakS5axaAbXgbxRH7cw3WC7IUEbQlNLD+oRHrNYLiwwHg8A==
Using [5] keys ok: 0Jlecf6ctgjx3RYujDsBon/zoJhcMU3u2aiuiVCaakS5axaAbXgbxRH7cw3WC7IUEbQlNLD+oRHrNYLiwwHg8A==
Using [6] keys ok: 0Jlecf6ctgjx3RYujDsBon/zoJhcMU3u2aiuiVCaakS5axaAbXgbxRH7cw3WC7IUEbQlNLD+oRHrNYLiwwHg8A==
```
