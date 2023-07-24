# Dice

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/manlikeNacho/Dice/blob/main/LICENSE)

A simple dice rolling application.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [License](#license)

## Introduction

Dice is a lightweight application that allows you to roll virtual dice. It can be used for various purposes such as board games. When a game is started, a commitment fee of 20 sats is deducted from the user’s wallet, and a number is generated randomly between 2 and 12 and stored. After a game has been started, the user can roll dice twice, and the result of both dice rolls is summed up. If the sum of the 2 numbers is equal to the number generated at game start-up, the user wins 10 sats.The user can roll dice as many times as he wants in a single game; however, each dice roll duo costs the user 5 sats and should be charged on the first die roll of the duo.

## Features

- Roll dice 
- Win or Lose game
- View transaction log
- Calculate the total sum of all rolls

## Installation

1. Clone the repository: git clone https://github.com/manlikeNacho/Dice.git
2. Navigate to the project directory: cd Dice
3. Install the necessary dependencies: go get

## Postman Collection
https://www.postman.com/lively-station-441669/workspace/school-adminstration/collection/21954692-0f47245d-548a-409e-aaea-5ae276ce8bab?action=share&creator=21954692

## License

This project is licensed under the [MIT License](https://github.com/manlikeNacho/Dice/blob/main/LICENSE).


