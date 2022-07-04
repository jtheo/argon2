# Argon2 Utility


## Help
```
Usage of bin/argon2-darwin-amd64:
  -encoded string
    	The encoded password to verify
  -iteration int
    	(timeCost)specifies the number of iterations of argon2. (default 4)
  -mem int
    	MemoryCost specifies the amount of memory to use in Kibibytes (default 3906)
  -paral int
    	Parallelism specifies the amount of threads to use. (default 1)
  -pass string
    	The password to encrypt
  -verbose
    	print some info about argon 2 parameters when generating a new password
  -version int
    	Version specifies the argon2 version to be used. (default 19)
```

## Example

### Generate Hashed Password
```
$: ./argon2 -pass mySuperPass
$argon2id$v=19$m=3906,t=4,p=1$ejcSAnKlTR6fhI0SRV7cUw$6EvE9S1qgsA8PoS96uZzuBGZZWCaeK91CHdHHYzn/mA
```

### Verify Hashed Password
```
$: ./argon2 -pass mySuperPass -encoded '$argon2id$v=19$m=3906,t=4,p=1$ejcSAnKlTR6fhI0SRV7cUw$6EvE9S1qgsA8PoS96uZzuBGZZWCaeK91CHdHHYzn/mA'
💚 Password match! 🔓

$: ./argon2 -pass mySuperPass -encoded '$argon2id$v=19$m=3906,t=4,p=1$ejcSAnKlTR6fhI0SRV7cUw$6EvE9S1qgsA8PoS96uZzuBGThis/is/wrong1hash000'
💥 Password don't match! 🔒
```