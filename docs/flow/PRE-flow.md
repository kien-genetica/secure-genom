```mermaid
sequenceDiagram
    participant A as Alice
    participant S as Server
    participant B as Bob
    participant C as Carlos

    Note over A: Generate keypair<br/>(aPriKey, aPubKey)
    Note over B: Generate keypair<br/>(bPriKey, bPubKey)
    Note over C: Generate keypair<br/>(cPriKey, cPubKey)

    %% Initial Encryption
    Note over A: Original message:<br/>"Hello, Proxy Re-Encryption"
    A->>A: Encrypt message with aPubKey<br/>Creates (cipherText, capsule)

    %% Re-encryption key generation for Bob
    A->>A: Generate re-encryption key (rkB)<br/>using (aPriKey, bPubKey)
    A->>S: Send rkB + capsule + cipherText

    %% Re-encryption key generation for Carlos
    A->>A: Generate re-encryption key (rkC)<br/>using (aPriKey, cPubKey)
    A->>S: Send rkC + capsule + cipherText

    %% Server Re-encryption for both recipients
    S->>S: Re-encrypt using rkB<br/>Creates newCapsuleB
    S->>S: Re-encrypt using rkC<br/>Creates newCapsuleC
    S->>B: Send newCapsuleB + cipherText
    S->>C: Send newCapsuleC + cipherText

    %% Decryption by both recipients
    B->>B: Decrypt using bPriKey<br/>Recovers original message
    C->>C: Decrypt using cPriKey<br/>Recovers original message

    Note over A: Can also decrypt original<br/>capsule with aPriKey
```
