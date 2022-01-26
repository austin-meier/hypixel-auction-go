Hypixel Auction
---
What is hypixel
Hypixel, officially the Hypixel Network, is a Minecraft minigame server released on April 13, 2013
> Starting out as a YouTube channel making Minecraft Adventure Maps, Hypixel is now one of the largest and highest quality Minecraft Server Networks in the world, featuring original games such as The Walls, Mega Walls, Blitz Survival Games, and many more!

The Problem?
---
Hypixel has a game mode based upon a variation of the SkyBlock map in Minecraft. This game has a live auction house in which players can list items to sell or browse and bid on other items players are auctioning. The Hypixel API provides this auction data but in an extremely undesirable paginated JSON output requiring you to iterate through multiple queries with no rules regarding duplicates or data formatting. This poses a large issue when attempting to parse the data, especially when speed is important such as in an auction house enviorment.

My goal is to utilize some of the fastest solutions modern languages offer to establish an efficient way to parse this data.

Why did I do this?
While I don't play Minecraft much anymore, at one point it still remains a quintessential part of my childhood and is a large contributor to my introduction and passion for programming.

One of my first real-world uses and projects with programming was creating Bukkit server plugins. This provided me a fun way to experience my creations as a useable feature and even provided an early source of income motivating me even further to continue growing as a developer.


Plan of Action
---
1. Web Framework
   - Fiber
2. Database
   - MonogDB
