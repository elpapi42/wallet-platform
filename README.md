# Bussiness Problem

We are engineers working in a Colombian fintech that is building the first iteration of its product from scratch. The core of the product is around a digital wallet (like Nequi or Daviplata) in which users will be able to deposit money to their wallets (Dollars, Euros, Bitcoins, Colombian Pesos, or any other currency), transfer that money to the wallets of other users in the platform, and finally, have the possibility of withdrawing the money to a bank account or cash.

As the first two engineers of this startup, is our responsibility to build the first iteration of the product, we are free to use any technologies, languages, or architecture we consider suitable to bring to life this new platform, but business stakeholders have few requirements we should meet to ensure the future of the platform:

1. We should have a platform with high tolerance to failures, so ideally, if something fails, our platform should be able to keep working, at least with reduced capabilities. This includes being able to recover quickly from those failures.
2. As we are dealing with money here, we can't accept that any kind of critical error generates inconsistencies with the deposit/transfer/withdraw as well as the balances of the wallets. We should create a platform that even in the worst conditions can be recoverable, with the balances and transactions as it should be, so the consistency of the platform under failures must never be compromised.

In short, we need a platform that gives priority to consistency and resiliency, over other things, like performance for example.

Also, we need to take into consideration that the business is operating with very little money, so they are not going to hire more engineers in at least the following two years. As we are only two engineers, we must ensure our design takes into consideration that we are going to maintain and evolve this for a long time, so we need a platform that can be maintained by two guys without trouble. We should also deliver the first iteration of the platform in the minimum possible time, but business stakeholders are willing to be patient if that means we can meet the requirements described above.

# Engineers Notes

1.

# Domain Draft

from typing import List
from uuid import UUID
from enum import Enum


class User:
    id: UUID
    email: str
    password: str
    name: str


class Currency(str, Enum):
    USD = "USD"
    EUR = "EUR"
    BTC = "BTC"
    COP = "COP"


class Wallet:
    id: UUID
    user_id: UUID
    currency: Currency
    balance: float

    def add_balance(self, amount: float):
        balance = self.balance + amount
    
        if balance < 0.0:
            raise ValueError('InsuficientFoundsError')

        self.balance = balance


class DepositSource:
    PSE = "PSE"
    BANK_TRANSFER = "BANK_TRANSFER"
    ATM_DEPOSIT = "ATM_DEPOSIT"
    CREDIT_CARD = "CREDIT_CARD"
    BITCOIN_NETWORK = "BITCOIN_NETWORK"

    def is_currency_supported(self, currency: Currency) -> bool:
        mapping = {
            DepositSource.PSE: [Currency.COP],
            DepositSource.BANK_TRANSFER: [Currency.COP],
            DepositSource.ATM_DEPOSIT: [Currency.COP],
            DepositSource.CREDIT_CARD: [Currency.COP, Currency.USD, Currency.EUR],
            DepositSource.BITCOIN_NETWORK: [Currency.BTC],
        }

        return currency in mapping[self]


class Deposit:
    id: UUID
    user_id: UUID
    source: DepositSource
    wallet: Wallet
    amount: float

    def validate_wallet_support_source(self):
        if not self.source.is_currency_supported(self.wallet.currency):
            raise ValueError('DepositSourceError')


class WithdrawalTarget:
    BANK_TRANSFER = "BANK_TRANSFER"
    ATM_WITHDRAWAL = "ATM_WITHDRAWAL"
    BITCOIN_NETWORK = "BITCOIN_NETWORK"

    def is_currency_supported(self, currency: Currency) -> bool:
        mapping = {
            WithdrawalTarget.BANK_TRANSFER: [Currency.COP],
            WithdrawalTarget.ATM_WITHDRAWAL: [Currency.COP],
            WithdrawalTarget.BITCOIN_NETWORK: [Currency.BTC],
        }

        return currency in mapping[self]

class Withdrawal:
    id: UUID
    user_id: UUID
    target: WithdrawalTarget
    wallet: Wallet
    amount: float

    def validate_wallet_support_target(self):
        if not self.target.is_currency_supported(self.wallet.currency):
            raise ValueError('WithdrawalTargetError')


class Transfer:
    id: UUID
    source_wallet: Wallet
    target_wallet: Wallet
    amount: float

    def validate_wallets_support_transfer(self):
        if not self.source_wallet.currency == self.target_wallet.currency:
            raise ValueError('TransferError')
