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
    subgraph User
        direction LR
        U[User]
    end

    subgraph Identity_Service
        direction LR
        IS[Identity Service]
    end

    subgraph Wallet_Service
        direction LR
        WS[Wallet Service]
    end

    subgraph Payment_Service
        direction LR
        PS[Payment Service]
    end

    subgraph Common_Package
        direction LR
    end

    U -- Sign Up/Login --> IS
    IS -- Publishes UserRegisteredEvent --> WS
    U -- Wallet Operations --> WS
    U -- Payment Operations --> PS
    PS -- gRPC Calls --> WS

    style IS fill:#bbf,stroke:#333,stroke-width:2px
    style WS fill:#bfb,stroke:#333,stroke-width:2px
    style PS fill:#f9f,stroke:#333,stroke-width:2px
    style U fill:#f66,stroke:#333,stroke-width:2px
