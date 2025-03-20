# Cribbage Challenge
A project to score a cribbage hand. The challenge came from the @cassidoo newsletter.

## Challenge Requirements
Write a function `scoreHand(cards)` that calculates the total score of a Cribbage hand. The input is an array of 5 card strings (including the starter card), where each card is represented as rank+suit (e.g., "AH", "10D", "KS"). Here are the scoring rules:
- 15s: 2 points for each combination of cards summing to 15
- Pairs: 2 points for each pair of same-rank cards
- Runs: 1 point per card in a run of 3 or more consecutive ranks

Example:
```
> scoreHand(["7H", "8C", "9D", "JH", "KS"])
> 5

> scoreHand(["AH", "2C", "3D", "4S", "5H"])
> 7
```

## To Run
This application is written in Golang. You can install Go on your machine using [https://golang.google.cn/dl/](https://golang.google.cn/dl/). Once installed, clone the project and use the following commands from the root directory of the project to see how things work.
```
// To run the tests
go test ./...

// To see the application in action
go run .
```
