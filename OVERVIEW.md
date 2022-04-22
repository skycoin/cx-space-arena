# cx-space-arena
## Summary of the Project
Space Arena consists a simulation environment where an agent is commanding a space ship within one given space structured on an Agent-Based Model. The commands the players are allowed to operate are: 
 * Move forward with the up arrow 
 * Turn left with the left arrow
 * Turn right with the right arrow
 * Spacebar lets the player to shoot bullets
 * CTRL key is a boost function which will propel the spaceship at a faster speed for a short period of time. There will be a boost bar at the bottom left corner of the screen which will refill over time once it is depleted. 
 * Shift key is a sonar ping that spreads across the arena allowing the sonar ping to detect other players and asteroids.

Surrounding the player will be various asteroids of different shapes and sizes. Depending on the size of the asteroid it will require the player to shoot more or less bullets to destroy the asteroid. Once a bigger asteroid is destroyed it will break off into smaller pieces, the number of pieces that the asteroid breaks off into is also dependent on the size of the initial asteroid.

The simulation will update on any time unit the agent makes a decision such as the positions of the agent, the positions of the bullets, the positions of the asteroids, the speed of the asteroid, and the rotation of the asteroid. 

The rate of fire coming from the spaceship depends on the tick rate. Assuming it the tick rate is 30 ticks per second, and the spaceship may only fire every 3 ticks, therefore the player will be limited to 10 shots or bullets per second. 

Each player ship will have a struct with different points such as data such as positions XY, velocity XY, player ID, direction in which the ship is turning. 

The sonar ping function will detect asteroids within their surroundings considering how far is the range of the ping for example 4 cm in diameter, then they can only see incoming asteroids within that diameter of the sonar ping. 

The agents will accumulate points depending on how many asteroids they destroy, each asteroid destroyed will initially count 15 points.
