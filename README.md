# Budget buddy

Budget Buddy is a user-friendly budget planner tool designed to simplify personal finance management. Users can effortlessly register and securely log in to the platform, where they can seamlessly add their income and expenses with detailed categorization with dates. The tool empowers users to track their cash flow efficiently, enabling them to manage both inflows and outflows effortlessly. With intuitive monthly and yearly reports, users gain valuable insights into their spending patterns through visual representations.

## Prerequisites

- latest version of go should be installed
- running mysql instance
- running redis instance

## Installation

- Clone project and open project directory

```bash
 git clone https://github.com/umer-emumba/BudgetBuddy
 cd BudgetBuddy
```

- install required depencies:

```
go mod tidy
```

- replace .env.example to .env and add required configurations
- build and run the project

```
 go run main.go
```
