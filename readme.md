# Secure Genome System Documentation

## System Overview

The Secure Genome System is a privacy-preserving platform for genetic data processing and sharing, utilizing Async Encryption, Proxy Re-Encryption (PRE) and Trusted Execution Environments (TEE) to ensure data security and controlled access.

## System Main flow

![System Main Flow](docs/flow/system-flow.md)

### Organization Flow

![Organization Flow](docs/flow/organization-flow.md)

## Core Components

### 1. Storage Service

- **Definition**:
  - A storage service is a node that stores the encrypted data. For example
    - a cloud storage service like AWS S3, Google Cloud Storage, Azure Blob Storage, etc.
    - Blockchain storage service like Filecoin, Arweave, etc.
    - Decentralized storage service like IPFS, etc.
- **Purpose**: Storage of encrypted data
- **Responsibilities**:
  - Store encrypted genetic data
  - Store encrypted re-encryption keys
  - Store encrypted processed results

### 2. Controller

- **Definition**:
  - A controller is our backend service that orchestrates the system.
- **Purpose**: System orchestration and access control
- **Responsibilities**:
  - System metadata management
  - Request routing
  - Data location tracking
  - Access permission verification
  - Payment management #TODO: can be move to another module

### 3. Validator

- **Definition**:
  - A validator is a trusted node that participates in the consensus process and validates the data. For example, a hospital or a research institution.
- **Purpose**: Validate the data and prepare the data for the consensus module.
- **Responsibilities**:
  - Encrypted data processing (fine-tuning, labeling, scoring)
  - Genetic data validation
  - Result labeling and scoring
  - Consensus preparation

### 4. Consensus Module (TEE)

- **Definition**:
  - A consensus module a module in our system that handles the consensus process.
- **Purpose**: Secure consensus, reward management for validators submitted results
- **Responsibilities**:
  - Validator result aggregation
  - Data validation / Risk score calculation #TODO: can define another module for this
  - Validator reward management
  - Secure data processing

### 5. Delegate Module (TEE)

- **Definition**:
  - A delegate module is a module in our system that handles the data sharing and access control.
- **Purpose**: Manage data sharing and access control
- **Responsibilities**:
  - Re-encryption key management
  - Proxy re-encryption operations
  - Organization access control
  - Data sharing management

## System Flows

### 1. PRE (Proxy Re-Encryption) setup flow

#### User - Delegate Module - Validator

1. User generates re-encryption keys for validators
2. Keys are encrypted for Delegate Module TEE
3. Encrypted keys are stored in Storage
4. Delegate Module TEE can decrypt the keys and perform re-encryption for validators

### 2. Organization Setup Flow

1. Organization registers through Controller
2. Controller verifies registration requirements
3. Consensus module generates re-encryption key
4. Key is transmitted to Delegate Module
5. Delegate Module stores C-O re-encryption key

### 3. PRE (Proxy Re-Encryption) Flow

1. Data owner (Alice) generates keypair
2. Data recipient (Bob) generates keypair
3. Alice encrypts message and creates capsule
4. Alice generates re-encryption key using private key and Bob's public key
5. Server performs re-encryption
6. Bob receives and decrypts data

## Security Model

The system employs multiple security mechanisms:

- Proxy Re-Encryption for secure data sharing
- Trusted Execution Environments for sensitive operations
- Access control at multiple levels
- Encrypted storage for all sensitive data
