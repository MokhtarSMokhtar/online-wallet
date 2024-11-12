# Online Wallet System

## Introduction

Welcome to the **Online Wallet System**, a microservices-based application designed to manage user authentication, wallet transactions, and payment processing. The system is composed of the following services:

1. **Identity Service**: Handles user registration, login, and authentication.
2. **Wallet Service**: Manages user wallet balances, transactions, coupon redemption, and fund transfers.
3. **Payment Service**: Processes wallet top-ups and payments, interacting with external payment gateways.

This README provides an overview of the system architecture, services, setup instructions, and usage guidelines.

## Architecture Diagram

```mermaid
graph LR
    subgraph "Identity Service"
        IdentityService[Identity Service]
    end

    subgraph "Wallet Service"
        WalletService[Wallet Service]
    end

    subgraph "Payment Service"
        PaymentService[Payment Service]
    end

    subgraph "Common Package"
        CommonPackage[Common Package]
    end

    User[User]

    User -->|Signup/Login| IdentityService
    IdentityService -->|Publishes Event| WalletService
    User -->|Wallet Operations| WalletService
    User -->|Payment Operations| PaymentService
    PaymentService -->|gRPC Calls| WalletService
    PaymentService -->|Uses| CommonPackage
    IdentityService -->|Uses| CommonPackage
    WalletService -->|Uses| CommonPackage

    style IdentityService fill:#bbf,stroke:#333,stroke-width:2px
    style WalletService fill:#bfb,stroke:#333,stroke-width:2px
    style PaymentService fill:#f9f,stroke:#333,stroke-width:2px
    style CommonPackage fill:#ff9,stroke:#333,stroke-width:2px
    style User fill:#f66,stroke:#333,stroke-width:2px
