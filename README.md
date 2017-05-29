# rex

```
$rex aes encrypt -k="key" -iv="iv" -p="encypt this" -m="ctr"
```

```
$rex aes encrypt -kf="/path/to/keyfile" -ivf="/path/to/iv" -pf="path/to/plaintext" -cf="path/to/ciphertext" -m="ctr"
```

## Encryption

| flag       | shortcut | description                                     | example                         |
|------------|----------|-------------------------------------------------|---------------------------------|
| key        | -k       | encryption key                                  | -k="this is a super secret key" |
| keyfile    | -kf      | read key from file                              | -kf="path/to/keyfile.rkf"       |
| iv         | -iv      | set the iv                                      | -iv="1234567890123456"          |
| ivfile     | -ivf     | read iv from file                               | -ivf="path/to/ivfile.riv"       |
| plain      | -p       | plaintext to encrypt                            | -p="encrypt this string"        |
| plainfile  | -pf      | read plaintext from file                        | -pf="path/to/plainfile.rpf"     |
| cipherfile | -cf      | file to output ciphertext, uses stdout if empty | -cf="path/to/output.rcf"        |

## Descryption 

| flag       | shortcut | description                                    | example                               |
|------------|----------|------------------------------------------------|---------------------------------------|
| key        | -k       | encryption key                                 | -k="this is a super secret key"       |
| keyfile    | -kf      | read key from file                             | -kf="path/to/keyfile.rkf"             |
| iv         | -iv      | set the iv                                     | -iv="1234567890123456"                |
| ivfile     | -ivf     | read iv from file                              | -ivf="path/to/ivfile.riv"             |
| ciphertext | -c       | ciphertext to decrypt                          | -c="imagine pseudo random bytes here" |
| cipherfile | -cf      | read ciphertext from file                      | -cf="path/to/cipherfile.rcf"          |
| plainfile  | -pf      | file to output plaintext, uses stdout if empty | -cf="path/to/output.rpf"              |
