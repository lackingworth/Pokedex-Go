## Installing
### Linux
* You must have [Go v1.24.3](https://go.dev/doc/install) (or higher) installed on your system
* Clone this repo to your desired location
* Run ```go build``` on your local machine in cloned repository
### Windows
* You can repeat the linux steps, it will work just fine
* Or alternatively you can [download an executable from release page](https://github.com/lackingworth/Pokedex-Go/releases)

## Info
 Pokédex in a command-line [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop). 
 It makes use of [PokéAPI](https://pokeapi.co/) to fetch all the data about Pokémon.
 A Pokédex is just a make-believe device that lets you look up information about Pokémon - things like their name, type, and stats.

## Available commands
* ```help``` - Displays all available commands in CLI
* ```catch <pokemon name>``` - Attempts to catch a Pokémon
> [!NOTE]  
> 
> You cannot catch the same Pokémon twice
> The probability of catching a Pokémon is dependent on their base experience (the higher the EXP the luckier you must be)
*  ```inspect <pokemon name>``` - Displays basic info about caught Pokémon
*  ```pokedex``` - Displays all caught Pokémon in this session
*  ```explore <location name>``` - Displays information about specified location
*  ```map``` - Displays location areas
> [!NOTE]  
> 
> First ```map``` command call displays the first 20 areas of the world.
> Each consecutive call will result in the next page (next 20 entries) shown
* ```mapb``` - Displays the previous page of location areas
* ```exit``` - Exits the Pokedex
* New features to come

## Version History

* v.0.0.1:

    * Initial Release
