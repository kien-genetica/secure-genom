```mermaid
sequenceDiagram
    participant U as User
    participant S as Storage
    participant O as Organization
    participant G as Controller
    participant V as Validator
    participant P as Processor (TEE)
    participant DM as Delegate Module (TEE)



    %% One-time P-DM-Organziation PRE Setup
    Note over O,DM: Organization Registration
    O->>G: Register Organization
    G->>G: Check requirements for registration
    G->>P: Register Re-encryption key for organization
    P->>P: Create Re-encryption key for organization (P-O)
    P->>DM: Send re-encryption key (P-O)
    DM->>DM: store P-O re-encryption key

    %% Re-encryption key setup
    Note over U,G: Re-encryption key setup
    G-->>U: List of un-authorized validators
    U->>U: 1. Generate re-encryption key (rk)<br/>for each validator <br/> 2. Encrypt RK for DM TEE
    U->>S: Store encrypted re-encryption keys
    U->>G: Mapping (user - validator RK) to re-encryption keys

    %% Data Upload with PRE
    Note over U,V: Data Upload Phase
    U->>U: Encrypt data with their own public key
    U->>S: Store encrypted data + public key
    U->>G: Submit data's id
    G->>G: Mapping file id to user

    %% Validator Processing
    Note over S,P: Validation Phase
    V->>G: Get data location/metadata
    G-->>V: Return metadata
    V->>S: Get data + re-encryption key (encrypted)
    S-->>V: Return data

    V->>DM: Request processing | data + re-encryption key (encrypted)
    Note over DM: perform proxy re-encryption
    DM->>DM: 1. Decrypt re-encryption key<br/>2. Perform re-encryption (Proxy-Re-encryption)
    DM->>V: Send re-encrypted data

    V->>V: 1. Decrypt with validator key<br/>2. Validate and label data <br/> 3. Encrypt for Consensus module

    %% Consensus Phase
    Note over P,S: Consensus / Validate Phase
    V->>P: Send processed data (P-encrypted)
    P->>P: 1. Process validations<br/>2. Calculate risk scores<br/>3. Handle rewards/slashing
    P->>P: Encrypt data for PRE
    P->>S: Store encrypted processed data
    P->>G: Submit result id

    %% Organization Access
    Note over S,G: Organization Access Phase
    O->>G: Request X records of case Y
    G->>G: 1. Check org permission <br/> 2. Check payment status
    G->>DM: Send corresponding record ids
    DM->>S: Get user encrypted data
    S-->>DM: return encrypted data (C-encrypted)

    Note over DM: proxy-re-encrypt for organization pubkey
    DM->>DM: 1. Fetch C-O re-encryption key <br/> 2. Re-Encrypt data for organization
    DM-->>O: Return re-encrypted data
    O->>O: Decrypt with org private key
```
