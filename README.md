# EDGE Toolkit
Multi purpose cross-platform cryptography tool for symmetric/asymmetric encryption, cipher-based message authentication code (CMAC), hash digest, hash-based message authentication code (HMAC), digital signature (ECDSA/EdDSA), shared key agreement (ECDH/X25519/VKO) and PBKDF2 function for embedded systems.
## Command-line ECC Integrated Security Suite
### Asymmetric Algorithms:
- Public Key Algorithms

    | Algorithm                   | 256 | 512 |ECDH |ECDSA|Encryption| TLS |
    |:----------------------------|:---:|:---:|:---:|:---:|:--------:|:---:|
    | Brainpool                   | O   | O   | O   | O   |          |     |
    | ECDSA (Secp256r1)           | O   |     | O   | O   |          | O   |
    | Ed25519/X25519              | O   |     | O   | O   |          | X   |
    | GOST R 34.10-2001 CryptoPro | O   |     | O   | O   |          |     |
    | GOST R 34.10-2012           | O   | O   | O   | O   |          | X   |
    | Koblitz (Secp256k1)         | O   |     | O   | O   | O        |     |
    | SM2                         | O   |     | O   | O   | O        | O   | 

    X: Windows 7 or higher.

### Symmetric Algorithms:
- Stream Ciphers:

    |      CIPHER      | KEY SIZE   |         MODE          |
    |:-----------------|:----------:|:---------------------:|
    | Chacha20Poly1305 |        256 | AEAD Stream Cipher    |
    | Salsa20          |        256 | -                     |
    | HC128            |        256 | -                     |
    | HC256            |        256 | -                     |

- Modes of operation:

    |MODE |                                | REQUIREMENTS                               |
    |:----|:-------------------------------|:-------------------------------------------|
    | CCM | Counter /w CBC-MAC (AEAD)      | Blocks of 128-bit with 128/192/256-bit keys|
    | EAX | Encrypt-Authenticate-Translate | Blocks of 128-bit with 128/192/256-bit keys|
    | GCM | Galois/Counter Mode (AEAD)     | Blocks of 128-bit with 128/192/256-bit keys|
    | MGM | Multilinear Galois Mode (AEAD) | Blocks of 64 or 128-bit with any key-length|
    | OCB | Offset Codebook Mode (AEAD)    | Blocks of 128-bit with 128/192/256-bit keys|
    | CTR | Counter Mode                   | All block ciphers                          |
    | OFB | Output Feedback Mode           | All block ciphers                          |

- 128-bit> Block Ciphers:

    |      CIPHER      | BLOCK SIZE |  KEY SIZE   |          MODES          |
    |:-----------------|:----------:|:-----------:|:-----------------------:|
    | Camellia         |        128 | 128/192/256 | All modes supported     |
    | Grasshopper      |        128 |         256 | All modes supported     |
    | RC6              |        128 | 128/192/256 | All modes supported     |
    | Rijndael (AES)   |        128 | 128/192/256 | All modes supported     |
    | SEED             |        128 |     128/256 | All modes supported     |
    | Serpent          |        128 | 128/192/256 | All modes supported     |
    | SM4              |        128 |         128 | All modes supported     |
    | Twofish          |        128 | 128/192/256 | All modes supported     |
    | Threefish256     |        256 |         256 | CTR and OFB             |
    | Threefish512     |        512 |         512 | CTR and OFB             |
    | Threefish1024    |       1024 |        1024 | CTR and OFB             |

- 64-bit block ciphers:

    |  CIPHER  | BLOCK SIZE |  KEY SIZE   |     MODES      |
    |:---------|:----------:|:-----------:|:--------------:|
    | 3DES     |          64|          192|MGM, CTR and OFB|
    | Blowfish |          64|  128/192/256|MGM, CTR and OFB|
    | CAST5    |          64|          128|MGM, CTR and OFB|
    | IDEA     |          64|          128|MGM, CTR and OFB|
    | GOST89 CryptoPro   |          64|          256|MGM, CTR and OFB|
    | HIGHT    |          64|          128|MGM, CTR and OFB|
    | Magma    |          64|          256|MGM, CTR and OFB|
    | RC5      |          64|          128|MGM, CTR and OFB|
    | Skipjack |          64|           80|MGM, CTR and OFB|
    | TEA      |          64|          128|MGM, CTR and OFB|
    | XTEA     |          64|          128|MGM, CTR and OFB|

- Hash Digests:

    |    ALGORITHM    | 128 | 160 | 192 | 256 | 512 | MAC |
    |:----------------|:---:|:---:|:---:|:---:|:---:|:---:|
    | BLAKE-2B        |     |     |     | O   | O   | O   |
    | BLAKE-2S        | O   |     |     | O   |     | O   |
    | GOST94 CryptoPro      |     |     |     | O   |     |     |
    | Grøstl          |     |     |     | O   |     |     |
    | JH              |     |     |     | O   |     |     |
    | Keccak          |     |     |     | O   | O   |     |
    | MD5 [Obsolete]  | O   |     |     |     |     |     |
    | Poly1305        | O   |     |     |     |     | O   |
    | RIPEMD          |     | O   |     |     |     |     |
    | SHA1 [Obsolete] |     | O   |     |     |     |     |
    | SHA2 (default)  |     |     |     | O   | O   |     | 
    | SHA3            |     |     |     | O   | O   |     |
    | SipHash         | O   |     |     |     |     | O   |
    | Skein512        |     |     |     | O   | O   | O   |
    | SM3             |     |     |     | O   |     |     |
    | Streebog        |     |     |     | O   | O   |     | 
    | Tiger           |     |     | O   |     |     |     | 
    | Whirlpool       |     |     |     |     | O   |     | 

### Cryptographic Functions:
  * Asymmetric Encryption
  * Symmetric Encryption
  * Digital Signature [ECDSA/EdDSA]
  * Shared Key Agreement [ECDH/X25519/VKO]
  * Recursive Hash Digest + Check
  * CMAC (Cipher-based message authentication code)
  * HMAC (Hash-based message authentication code)
  * PBKDF2 (Password-based key derivation function 2)
  * TLS (Transport Layer Security) with SM2, ECDSA, Ed25519 and GOST

### Non-Cryptographic Functions:
  * Base64 string conversion
  * Bin to Hex/Hex to Bin string conversion
  * Data sanitization method
  * DJB2, DJB2a, SDMB, ELF32 hash functions (CRC32 equivalent)
  * LZMA, XZ and GZIP compression
  * Random Art Public key Fingerprint (ssh-keygen equivalent)

## Usage:
<pre> -algorithm string
       Asymmetric algorithm: brainpool256r1, ecdsa, sm2. (default "ecdsa")
 -bits int
       Key length: 64, 128, 192 or 256. (for RAND and PBKDF2) (default 256)
 -check string
       Check hashsum file. (- for STDIN)
 -cipher string
       Symmetric algorithm, e.g. aes, serpent, twofish. (default "aes")
 -crypt string
       Encrypt/Decrypt with symmetric block ciphers.
 -derive
       Derive shared secret key.
 -digest string
       Target file/wildcard to generate hashsum list. (- for STDIN)
 -iter int
       Iterations. (for PBKDF2 and SHRED commands) (default 1)
 -key string
       Private/Public key, password or HMAC key, depending on operation.
 -keygen
       Generate asymmetric keypair.
 -list
       List all available algorithms.
 -mac string
       Compute Cipher-based/Hash-based message authentication code.
 -md string
       Hash algorithm, e.g. sha256, sm3 or keccak256. (default "sha256")
 -mode string
       Mode of operation: CCM, GCM, MGM, OCB, CTR or OFB. (default "CTR")
 -pbkdf2
       Password-based key derivation function.
 -pkeyutl string
       Encrypt/Decrypt with asymmetric algorithms secp256k1 and sm2.
 -pub string
       Remote's side public key. (for shared secret derivation)
 -rand
       Generate random cryptographic key.
 -recursive
       Process directories recursively. (for DIGEST command only)
 -salt string
       Salt. (for PBKDF2 only)
 -shred string
       Target file/path/wildcard to apply data sanitization method.
 -sign
       Sign hash with Private key.
 -signature string
       Input signature. (verification only)
 -tcp string
       Encrypted TCP/IP Transfer Protocol. [dump|ip|send]
 -util string
       Utilities for encoding and compression. (type -util help)
 -verbose
       Verbose mode. (for CHECK, the exit code is always 0 in this mode)
 -verify
       Verify signature with Public key.</pre>
## Examples:
### Asymmetric keypair generation (ECDSA):
```sh
./edgetk -keygen 
```
### Symmetric key generation (default 256):
```sh
./edgetk -rand [-bits 64|128|192|256]
```
### Digital signature:
```sh
./edgetk -sign -key $prvkey < file.ext > sign.txt
sign=$(cat sign.txt)
./edgetk -verify -key $pubkey -signature $sign < file.ext
```
### Shared key agreement (ECDH/X25519/VKO):
```sh
./edgetk -derive -key $prvkey -pub $pubkey
```
### Encryption/decryption with asymmetric cipher SM2 or Secp256k1:
```sh
./edgetk -pkeyutl enc -key $pubkey -algorithm sm2 < plaintext.ext 
./edgetk -pkeyutl dec -key $prvkey -algorithm sm2 < ciphertext.ext 
```
### Encryption/decryption with symmetric block cipher:
```sh
./edgetk -crypt enc -key $256bitkey < plaintext.ext > ciphertext.ext
./edgetk -crypt dec -key $256bitkey < ciphertext.ext > plaintext.ext
```
### SHA256-based HMAC:
```sh
./edgetk -mac hmac -key $256bitkey < file.ext
```
### PBKDF2 (password-based key derivation function 2):
```sh
./edgetk -pbkdf2 -key "pass" -iter 10000 -salt "salt"
```
### Note:
The PBKDF2 function can be combined with the CRYPT and HMAC commands:
```sh
./edgetk -crypt enc -pbkdf2 -key "pass" < plaintext.ext > ciphertext.ext
./edgetk -mac hmac -pbkdf2 -key "pass" -iter 10000 -salt "salt" < file.ext
```

### Shred (data sanitization method, 25 iterations):
```sh
./edgetk -shred keypair.ini -iter 25
```
### Bin to Hex/Hex to Bin:
```sh
echo somestring|./edgetk -util hexenc
echo hexstring|./edgetk -util hexdec
```
### Base64 encoding:
```sh
echo somestring|./edgetk -util b64enc
echo b64string|./edgetk -util b64dec
```
### String compression LZMA/XZ/GZIP:
```sh
echo somestring|./edgetk -util compress -algorithm lzma
type lzmafile|./edgetk -util decompress -algorithm lzma
```
### TCP/IP Dump/Send:
```sh
./edgetk -tcp ip > PublicIP.txt
./edgetk -tcp dump [-pub "8081"] > Pubkey.txt
./edgetk -tcp send [-pub "127.0.0.1:8081"] < Pubkey.txt
```
### Random Art (Public Key Fingerprint):
```sh
./edgetk -key $pubkey
./edgetk -key - < Pubkey.txt
```
## License

This project is licensed under the ISC License.
#### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Research Lab.
