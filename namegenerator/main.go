package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	typeFlag = flag.String("type", "card", "do you want a pokemon or card name?")

	// Ranks are the card ranks in a normal deck of cards
	Ranks = []string{
		"two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"ace", "king", "queen", "jack",
	}

	// Suits are the card suits in a normal deck of cards
	Suits = []string{
		"clubs", "hearts", "spades", "diamonds",
	}

	// Names is an array of the English names of all Pokemon in the Kanto region
	Names = []string{
		"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon",
		"Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie",
		"Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill",
		"Pidgey", "Pidgeotto", "Pidgeot", "Rattata", "Raticate",
		"Spearow", "Fearow", "Ekans", "Arbok", "Pikachu",
		"Raichu", "Sandshrew", "Sandslash", "Nidoran♀", "Nidorina",
		"Nidoqueen", "Nidoran♂", "Nidorino", "Nidoking", "Clefairy",
		"Clefable", "Vulpix", "Ninetales", "Jigglypuff", "Wigglytuff",
		"Zubat", "Golbat", "Oddish", "Gloom", "Vileplume",
		"Paras", "Parasect", "Venonat", "Venomoth", "Diglett",
		"Dugtrio", "Meowth", "Persian", "Psyduck", "Golduck",
		"Mankey", "Primeape", "Growlithe", "Arcanine", "Poliwag",
		"Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam",
		"Machop", "Machoke", "Machamp", "Bellsprout", "Weepinbell",
		"Victreebel", "Tentacool", "Tentacruel", "Geodude", "Graveler",
		"Golem", "Ponyta", "Rapidash", "Slowpoke", "Slowbro",
		"Magnemite", "Magneton", "Farfetch'd", "Doduo", "Dodrio",
		"Seel", "Dewgong", "Grimer", "Muk", "Shellder",
		"Cloyster", "Gastly", "Haunter", "Gengar", "Onix",
		"Drowzee", "Hypno", "Krabby", "Kingler", "Voltorb",
		"Electrode", "Exeggcute", "Exeggutor", "Cubone", "Marowak",
		"Hitmonlee", "Hitmonchan", "Lickitung", "Koffing", "Weezing",
		"Rhyhorn", "Rhydon", "Chansey", "Tangela", "Kangaskhan",
		"Horsea", "Seadra", "Goldeen", "Seaking", "Staryu",
		"Starmie", "Mr. Mime", "Scyther", "Jynx", "Electabuzz",
		"Magmar", "Pinsir", "Tauros", "Magikarp", "Gyarados",
		"Lapras", "Ditto", "Eevee", "Vaporeon", "Jolteon",
		"Flareon", "Porygon", "Omanyte", "Omastar", "Kabuto",
		"Kabutops", "Aerodactyl", "Snorlax", "Articuno", "Zapdos",
		"Moltres", "Dratini", "Dragonair", "Dragonite", "Mewtwo",
		"Mew",
	}
	// Moves is every single move from the original Pokemon game.
	Moves = []string{
		"Pound", "KarateChop", "DoubleSlap", "CometPunch", "MegaPunch", "PayDay",
		"FirePunch", "IcePunch", "ThunderPunch", "Scratch", "ViceGrip", "Guillotine",
		"RazorWind", "SwordsDance", "Cut", "Gust", "WingAttack", "Whirlwind",
		"Fly", "Bind", "Slam", "VineWhip", "Stomp", "DoubleKick",
		"MegaKick", "JumpKick", "RollingKick", "SandAttack", "Headbutt", "HornAttack",
		"FuryAttack", "HornDrill", "Tackle", "BodySlam", "Wrap", "TakeDown",
		"Thrash", "Double-Edge", "TailWhip", "PoisonSting", "Twineedle", "PinMissile",
		"Leer", "Bite", "Growl", "Roar", "Sing", "Supersonic",
		"SonicBoom", "Disable", "Acid", "Ember", "Flamethrower", "Mist",
		"WaterGun", "HydroPump", "Surf", "IceBeam", "Blizzard", "Psybeam",
		"BubbleBeam", "AuroraBeam", "HyperBeam", "Peck", "DrillPeck", "Submission",
		"LowKick", "Counter", "SeismicToss", "Strength", "Absorb", "MegaDrain",
		"LeechSeed", "Growth", "RazorLeaf", "SolarBeam", "PoisonPowder", "StunSpore",
		"SleepPowder", "PetalDance", "StringShot", "DragonRage", "FireSpin", "ThunderShock",
		"Thunderbolt", "ThunderWave", "Thunder", "RockThrow", "Earthquake", "Fissure",
		"Dig", "Toxic", "Confusion", "Psychic", "Hypnosis", "Meditate",
		"Agility", "QuickAttack", "Rage", "Teleport", "NightShade", "Mimic",
		"Screech", "DoubleTeam", "Recover", "Harden", "Minimize", "Smokescreen",
		"ConfuseRay", "Withdraw", "DefenseCurl", "Barrier", "LightScreen", "Haze",
		"Reflect", "FocusEnergy", "Bide", "Metronome", "MirrorMove", "Self-Destruct",
		"EggBomb", "Lick", "Smog", "Sludge", "BoneClub", "FireBlast",
		"Waterfall", "Clamp", "Swift", "SkullBash", "SpikeCannon", "Constrict",
		"Amnesia", "Kinesis", "Soft-Boiled", "HighJumpKick", "Glare", "DreamEater",
		"PoisonGas", "Barrage", "LeechLife", "LovelyKiss", "SkyAttack", "Transform",
		"Bubble", "DizzyPunch", "Spore", "Flash", "Psywave", "Splash",
		"AcidArmor", "Crabhammer", "Explosion", "FurySwipes", "Bonemerang", "Rest",
		"RockSlide", "HyperFang", "Sharpen", "Conversion", "TriAttack", "SuperFang",
		"Slash", "Substitute", "Struggle",
	}
)

// Pokemon generates a new name based on the moves and Pokemon names
func Pokemon() string {
	move1 := Moves[rand.Int()%len(Moves)]
	move2 := Moves[rand.Int()%len(Moves)]
	name := Names[rand.Int()%len(Names)]

	return strings.ToLower(fmt.Sprintf("%s-%s-%s", move1, move2, name))
}

// Card creates a name.
func Card() string {
	rank := Ranks[rand.Int()%len(Ranks)]
	suit := Suits[rand.Int()%len(Suits)]

	return fmt.Sprintf("%s-of-%s-%d", rank, suit, rand.Int63()%100000)
}

func main() {
	flag.Parse()

	if *typeFlag == "pokemon" {
		fmt.Println(Pokemon())
	} else {
		fmt.Println(Card())
	}
}
