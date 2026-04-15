### Factory Design Pattern: Write Once, Extend Forever

You know that time when you have a very nice if-else block and your boss tells you, Hey, can we add one more option?

The one sentence typically implies that you need to break open your clean code, push another condition into an already growing monster and hope that you did not mess with the logic that had already been in place to accommodate the ten previously existing choices.

And there is a way that is better. It is called the Factory Pattern.

This is a post in which we will develop a Payment Gateway System in Go. We will begin with the messy way and understand the reason why it hurts followed by a fix to a pattern that will enable you to add bKash, Nagad, or even SpaceBucks to your files as easily as you would add a new file.

## The Issue: The vicious never-ending if-else.

Suppose that you are developing an e-commerce site. Users are able to pay through Credit Card, PayPal or Bank Transfer. Your code is likely to resemble the following:

```go
func processPayment(paymentType string, amount float64) {
    if paymentType == "creditcard" {
        // 50 lines of credit card logic.
    } else if paymentType == "paypal" {
        // 40 lines of PayPal magic.
    } else if paymentType == "banktransfer" {
        // 30 lines of bank code.
    }
}
```

This works. Until Monday morning you should add Crypto. Then bKash. Then Nagad.

This strategy is fatally flawed in three ways:

1. Code Bloat: Your code turns into an if-else novel.
2. The OCP Violation: You continue changing old code to implement new features. New bugs are introduced by touching old code.
3. Testing Nightmare: You can not test PayPal alone, without dragging the entire functionality down.

## The Solution: Factory Pattern.

There is only one rule of the Factory Pattern and this is: Do not call the builder directly. Get one made by a Factory.

Real-Life example: The Coffee Shop.

You walk into a coffee shop. You do not go behind the counter, grab the espresso machine, and start steaming milk. You say: "One cappuccino, please."

The barista (The Factory) decides how to make it. You just receive a Coffee (the interface with caffeine and flavor) and you Drink().

Now replace "Coffee" with "Payment" and "Drink()" with "Pay()". Same concept.

In code language: Your primary application must state: "Give me a Payment Processor." It does not matter whether it is a CreditCard or PayPal struct under the hood. It just wants to call .Pay().

## Step 1: The Contract (The Interface)

We have to have a contract before we construct anything. In Go, we make use of an interface. This interface writes: "To be a Payment Method in this system you have to understand how to Pay() or not."

Create a file: payment/payment.go

```go
package payment

// Our contract is payment. Whoever does this can process money.
type Payment interface {
    Pay(amount float64) error
}
```

Analogy: This is the electric socket in your wall. What you plug in (Fan, TV, Fridge) is unknown to it. It does not mind whether the shape (the method signature) will take the plug. The socket is powered, as long as it fits.

Think of it like a coffee shop menu. The menu (interface) says: "Every drink we serve must be sippable." Whether it is an Espresso, Cappuccino, or Latte, they all know how to be sipped. You do not care how the barista made it. You just Sip().

## Step 2: The Concrete Builders (CreditCard, PayPal, etc.)

We now make the real workers. All of them will sign the contract by implementing the Payment interface.

Write a file: payment/types.go.

```go
package payment

import "fmt"

type CreditCard struct{}
func (c *CreditCard) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using Credit Card (Bank deducting now...)\n", amount)
    return nil
}

type PayPal struct{}
func (p *PayPal) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using PayPal (Login successful!)\n", amount)
    return nil
}

type BankTransfer struct{}
func (b *BankTransfer) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using Bank Transfer (Check your IBAN!)\n", amount)
    return nil
}
```

Notice: In Go we do not write implements Payment. Go is smart. Since these structs include a Pay(amount float64) error method, Go automatically understands them to be Payment types. This is referred to as implicit interface implementation.

Back to our coffee shop: CreditCard is like an Espresso. PayPal is like a Cappuccino. BankTransfer is like a Latte. They are made differently behind the counter, but they all come in a cup and they all satisfy your caffeine craving.

## Step 3: Factory itself.

Here the magic occurs. This function takes a string like "paypal" and gives you back a Payment Interface.

Make a file: payment/factory.go.

```go
package payment

import "fmt"

// PaymentFactory decides who to hire for the job.
func PaymentFactory(paymentType string) (Payment, error) {
    switch paymentType {
    case "creditcard":
        return &CreditCard{}, nil
    case "paypal":
        return &PayPal{}, nil
    case "banktransfer":
        return &BankTransfer{}, nil
    default:
        return nil, fmt.Errorf("Sorry, we do not support %s yet", paymentType)
    }
}
```

This is the barista. You tell the barista "cappuccino" and they hand you back a Coffee. You do not need to know how they steamed the milk or pulled the espresso shot. You just receive something that satisfies the Coffee interface.

## Step 4: Try It in main.go.

Look now how clear our main argument is. It has nothing to do with Credit Cards or PayPal. It is merely aware of the Contract.

```go
package main

import (
    "fmt"
    "payment-factory/payment"
)

func main() {
    // We want to test all our methods
    methods := []string{"creditcard", "paypal", "banktransfer", "bkash"}

    for _, method := range methods {
        fmt.Printf("\nTrying to pay via: %s\n", method)

        // Ask the factory for a PaymentProcessor
        PaymentProcessor, err := payment.PaymentFactory(method)

        if err != nil {
            fmt.Println(err)
            continue // Skip if method not found
        }

        // We just call Pay(). We do not know the concrete type.
        PaymentProcessor.Pay(500.00)
    }
}
```

This is you at the coffee shop counter. You do not need to know how to make a cappuccino. You just say the word, receive the cup, and drink.

## Step 5: The Superpower (Extending Without Breaking)

This is the actual payoff. We would like to include bKash support.

What is it we should change?

- payment/types.go: Create a new struct BKash.
- payment/factory.go: Add a new case, bkash.
- main.go: Do not touch. Not even one of the characters.

New code in types.go:

```go
type BKash struct{}
func (b *BKash) Pay(amount float64) error {
    fmt.Printf("Paid %.2f using bKash (PIN verified!)\n", amount)
    return nil
}
```

New case in factory.go:

```go
case "bkash":
    return &BKash{}, nil
```

Run it again. The loop works well with bkash. The main.go file is Closed to Modification, but Open to Extension. Here is the holy grail of maintainable code.

This is like your coffee shop adding a Mocha to the menu. The barista learns a new recipe (new case in factory). The menu board gets updated (new struct in types.go). But you, the customer, still walk up to the counter and say "One Mocha, please." Your ordering process does not change.

![Factory Design Pattern](factory-design-pattern/payment-gateway-go/factory-pattern.png)

## And one more thing: What is PaymentProcessor Really?

When developing, you may peep under the hood with the reflect package. You may wonder: Why does reflect show \*payment.CreditCard in case the Factory is returning a Payment interface?

This is a great question, and knowing it elevates you to a better Go developer.

Consider the PaymentProcessor variable to be a Coffee Cup.

- Static Type (The Cup): The outside of the cup says "Coffee" on it. This is what the compiler checks at compile time. You can only Drink() from it.
- Dynamic Type (The Liquid Inside): The actual liquid inside the cup is Espresso or Cappuccino. This is what reflect reveals at runtime.
- Kind (The Pointer): ptr (Pointer). Since we returned &CreditCard, the liquid is served at a specific temperature.

The following is the code snippet that will testify:

```go
PaymentProcessor, _ := payment.PaymentFactory("creditcard")
fmt.Printf("Type: %T\n", PaymentProcessor) // Output: *payment.CreditCard (The liquid inside!)
```

The Factory pattern is good since the caller never touches anything other than The Cup (Payment). It is The Barista (Factory) who is allowed to pour The Liquid (\*CreditCard). It is this decoupling that makes your code so flexible.

Next time your boss asks for a new payment method, you can confidently say: "Sure, give me two minutes." And you will not even break a sweat. Just like adding a new syrup to the coffee bar.
