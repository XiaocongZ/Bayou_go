# Bayou System: A Shoot Out in the Darkness
This project for Distributed Algorithms build upon the Bayou architecture a textual game.

It simulates a shoot out in the darkness(1 vs 1), where you only see the opponent when he fires. Press 'a' fro moving left, press 'd' for moving right, press 'f' for firing a shot. Firing exposes one's position.

### how to build
`make all`

### how to run
To run the game, with a primary server, a local server&client with random AI, a local server&client for control:

Terminal 1:
`./bayou_primary`

Terminal 2:
`./bayou_bot`

Terminal 3:
`./bayou_server`
