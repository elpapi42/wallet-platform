# Bussiness Problem

We are engineers working in a Colombian fintech that is building the first iteration of its product from scratch. The core of the product is around a digital wallet (like Nequi or Daviplata) in which users will be able to deposit money to their wallets (Dollars, Euros, Bitcoins, Colombian Pesos, or any other currency), transfer that money to the wallets of other users in the platform, and finally, have the possibility of withdrawing the money to a bank account or cash.

As the first two engineers of this startup, is our responsibility to build the first iteration of the product, we are free to use any technologies, languages, or architecture we consider suitable to bring to life this new platform, but business stakeholders have few requirements we should meet to ensure the future of the platform:

1. We should have a platform with high tolerance to failures, so ideally, if something fails, our platform should be able to keep working, at least with reduced capabilities. This includes being able to recover quickly from those failures.
2. As we are dealing with money here, we can't accept that any kind of critical error generates inconsistencies with the deposit/transfer/withdraw as well as the balances of the wallets. We should create a platform that even in the worst conditions can be recoverable, with the balances and transactions as it should be, so the consistency of the platform under failures must never be compromised.

In short, we need a platform that gives priority to consistency and resiliency, over other things, like performance for example.

Also, we need to take into consideration that the business is operating with very little money, so they are not going to hire more engineers in at least the following two years. As we are only two engineers, we must ensure our design takes into consideration that we are going to maintain and evolve this for a long time, so we need a platform that can be maintained by two guys without trouble. We should also deliver the first iteration of the platform in the minimum possible time, but business stakeholders are willing to be patient if that means we can meet the requirements described above.

# Engineers Notes

1.
