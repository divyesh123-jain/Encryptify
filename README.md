# Encryptify

Encryptify is a simple command-line tool for file encryption and decryption using the AES encryption algorithm, built with GoLang. With Encryptify, you can easily encrypt your sensitive files to protect them from unauthorized access and decrypt them when needed.

## Features

- Encrypt files securely using AES encryption.
- Decrypt encrypted files to recover the original content.
- Easy-to-use command-line interface.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/encryptify.git
    ```

2. Build the executable:

    ```bash
    cd encryptify
    go build
    ```

3. Run Encryptify using the generated executable:

    - On Linux/macOS:

    ```bash
    ./encryptify [options]
    ```

    - On Windows:

    ```bash
    .\encryptify.exe [options]
    ```

## Usage

Encryptify supports the following command-line options:

- `-mode`: Specifies the mode of operation (`encrypt` or `decrypt`).
- `-input`: Specifies the input file path.
- `-output`: Specifies the output file path.
- `-key`: Specifies the encryption/decryption key (16, 24, or 32 bytes).

### Encrypt a file:
```bash
./encryptify -mode=encrypt -input=input.txt -output=output.enc -key=yourencryptionkey
```

## Decrypt a file:
```bash
./encryptify -mode=decrypt -input=output.enc -output=output.txt -key=yourencryptionkey
```
## Contributing

Contributions are welcome! If you have suggestions, bug reports, or feature requests, please open an issue or submit a pull request

