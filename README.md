# open-payment-gateway

Open payment gateway(OPG) is a self-hosted tool to help you receive your payments in crypto currencies without relying on any 3rd party crypto payment processors. Yayy, no processing fees anymore!

Right now we're supporting network's native currency transactions on all evm compatible blockchains, such as:  Ethereum, Binance Smart Chain, Polygon, Arbitrum, etc.
ERC20 token transfers will be available in the next releases.

OPG supports HD wallets, so no need to backup addresses periodically in order to protect your assets.
All of your addresses are recoverable using a master public key(xpub). NO NEED TO PUT YOUR PRIVATE KEYS ONLINE to create addresses.

This increase privacy, since even in case of having an uncontrolled access to your server, no one can move your funds to another address because there are no private keys saved in the server, tho they can monitor all of your addresses since you've saved the master public key in the server.

OPG does have a few built-in tools for managing your funds, for example calculating the total balance of all the addresses, signing transactions and moving funds to another address.
More tools will be available in the next releases.

## Hardware wallet support will be added in v0.8

Take a look at [[roadmap section]].

OPG was written to be used with minimum configuration.
OPG supports 3rd party web3 providers, so you don't necessarily need any self-hosted node to operate. This significantly reduces your costs for 2 main reasons:

1. Running a personal node may require a huge amount of CPU, RAM, and storage resources, hence it would increase your server costs.

2. Sometimes maintaining a personal node is a pain in the back, there are some days that the server crashes, you run out of storage, network is down, and all the other scenarios, and all of them require you to manually look into the problem and monitor the nodes all the time. Why don't we delegate this task to someone else, like a third party web3 provider for a cheap price or sometimes free?

It does not mean that you can't connect it to your own nodes.

The documentation is still being developed, so please be patient.

Contribution is welcomed from everyone

## Funding and Development Goals

Open Payment Gateway (OPG) relies on the generous support of the community to continue its development and expansion. Your contributions, in the form of Bitcoin donations, play a pivotal role in helping us reach our funding goals. These goals, when achieved, enable us to introduce new features and improvements to OPG, benefiting all users. Here's how it works:

## Donations

To support OPG's development, you can make a Bitcoin donation to our project. Your donation helps us cover development costs, maintain the project infrastructure, and allocate more time and resources to enhance OPG's capabilities. Every contribution, regardless of its size, is greatly appreciated.

**Bitcoin Address for Donations**: `A bitcoin Address`

We understand that not everyone can contribute through sponsorship, so donations provide an alternative way to support our project's growth.

## Funding Goals

We've set specific funding goals that are tied to the total amount of donations received. These goals serve as milestones for the project's development. As we reach each funding goal, we commit to implementing new features and improvements and making them available to all OPG users. Here are our current funding goals:

1. **Goal 1: [Goal Name]**
   - **Total Donation Target**: [Bitcoin Amount]
   - **Features to be Implemented**: [List of features or improvements]

2. **Goal 2: [Goal Name]**
   - **Total Donation Target**: [Bitcoin Amount]
   - **Features to be Implemented**: [List of features or improvements]

3. **Goal 3: [Goal Name]**
   - **Total Donation Target**: [Bitcoin Amount]
   - **Features to be Implemented**: [List of features or improvements]

## Progress and Updates

We will regularly update the progress towards our funding goals in this README, ensuring transparency about the state of the project. You can also check the latest status and announcements on our [project's website](link_to_project_website).

## Thank You for Your Support

Your contributions, whether through donations or PR submission, are essential in making OPG a feature-rich and reliable self-hosted payment gateway. We greatly appreciate your support, and together, we'll make OPG even more powerful and user-friendly for the community.

Please consider donating and helping us achieve our funding goals. Together, we can shape the future of Open Payment Gateway!
