# Features

## MVP
- type a screen name when you open the site and choose yellow or red
- one person pick yellow; the other pick red
- can play a game and see a winner and restart
- tablet first, should run on mobile and PC

## shared view
- tap and drag the piece to place
- player can see what the other player is selecting

## authentication
- login, creates a permanent profile
- add friends

## finding a match
- join next available game
- by MMR
- lobby

## multiplayer
- two or more people can join a team
- alternative moves

## game rules
- limited time per move; if you miss, game picks randomly
- pause and resart game
- mega board

## AI
- make a smart AI ???

# Tech
HTMX and Golang - yes

HTMX and Fastapi - no because I don't like Python
HTMX and C# - no because HTML templating in C# looks annoying
Angular and C# - no because it will be so much bloat

CSS for chip transitions like this: https://github.com/Kamide/connect-n/tree/main

## Design
- each game generates an ID which goes in the URL
- we save the game state with every move, so at any time the user can refresh and see latest position
- web sockets (?) for live updates of the other player's move
