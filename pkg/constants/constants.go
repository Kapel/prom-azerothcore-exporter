package constants

import "fmt"

// WoW race to faction mapping
var RaceToFaction = map[int]string{
	1:  "Alliance", // Human
	2:  "Horde",    // Orc
	3:  "Alliance", // Dwarf
	4:  "Alliance", // Night Elf
	5:  "Horde",    // Undead
	6:  "Horde",    // Tauren
	7:  "Alliance", // Gnome
	8:  "Horde",    // Troll
	9:  "Horde",    // Goblin
	10: "Horde",    // Blood Elf
	11: "Alliance", // Draenei
	22: "Horde",    // Worgen
	24: "Alliance", // Pandaren (Neutral)
	25: "Alliance", // Pandaren (Alliance)
	26: "Horde",    // Pandaren (Horde)
}

// WoW race names
var RaceNames = map[int]string{
	1:  "Human",
	2:  "Orc",
	3:  "Dwarf",
	4:  "Night Elf",
	5:  "Undead",
	6:  "Tauren",
	7:  "Gnome",
	8:  "Troll",
	9:  "Goblin",
	10: "Blood Elf",
	11: "Draenei",
	22: "Worgen",
	24: "Pandaren (Neutral)",
	25: "Pandaren (Alliance)",
	26: "Pandaren (Horde)",
}

// WoW class names
var ClassNames = map[int]string{
	1:  "Warrior",
	2:  "Paladin",
	3:  "Hunter",
	4:  "Rogue",
	5:  "Priest",
	6:  "Death Knight",
	7:  "Shaman",
	8:  "Mage",
	9:  "Warlock",
	10: "Monk",
	11: "Druid",
	12: "Demon Hunter",
}

// Helper functions for readable names
func GetDifficultyName(difficulty int) string {
	difficultyNames := map[int]string{
		0: "Normal",
		1: "Heroic",
		2: "10_Player",
		3: "25_Player",
		4: "10_Player_Heroic",
		5: "25_Player_Heroic",
	}
	if name, exists := difficultyNames[difficulty]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", difficulty)
}

func GetLFGStateName(state int) string {
	stateNames := map[int]string{
		0: "None",
		1: "RoleCheck",
		2: "Queued",
		3: "Proposal",
		4: "Boot",
		5: "Dungeon",
		6: "FinishedDungeon",
	}
	if name, exists := stateNames[state]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", state)
}

func GetIPActionTypeName(actionType int) string {
	typeNames := map[int]string{
		0: "Login",
		1: "Failed_Login",
		2: "Logout",
		3: "Character_Create",
		4: "Character_Delete",
		5: "Character_Login",
		6: "Character_Logout",
		7: "Password_Change",
		8: "Account_Create",
		9: "Account_Delete",
	}
	if name, exists := typeNames[actionType]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", actionType)
}

func GetLagTypeName(lagType int) string {
	lagTypeNames := map[int]string{
		0: "World",
		1: "Instance",
		2: "Battleground",
		3: "Arena",
		4: "Raid",
	}
	if name, exists := lagTypeNames[lagType]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", lagType)
}

func GetBattlegroundTypeName(bgType int) string {
	bgTypeNames := map[int]string{
		1:  "Alterac Valley",
		2:  "Warsong Gulch",
		3:  "Arathi Basin",
		4:  "Eye of the Storm",
		5:  "Strand of the Ancients",
		6:  "Isle of Conquest",
		7:  "Twin Peaks",
		8:  "Battle for Gilneas",
		9:  "Temple of Kotmogu",
		10: "Silvershard Mines",
		11: "Deepwind Gorge",
	}
	if name, exists := bgTypeNames[bgType]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", bgType)
}

var MapNames = map[int]string{
	0:   "Eastern Kingdoms",
	1:   "Kalimdor",
	530: "Outland",
	571: "Northrend",
	609: "Ebon Hold",
	30:  "Alterac Valley",
	489: "Warsong Gulch",
	529: "Arathi Basin",
	566: "Eye of the Storm",
	607: "Strand of the Ancients",
	628: "Isle of Conquest",
}

var MapToContinent = map[int]string{
	0:   "Eastern Kingdoms",
	1:   "Kalimdor",
	530: "Outland",
	571: "Northrend",
	609: "Ebon Hold",
}

func GetMapName(mapID int) string {
	if name, ok := MapNames[mapID]; ok {
		return name
	}
	return fmt.Sprintf("Map_%d", mapID)
}

func GetContinentName(mapID int) string {
	if name, ok := MapToContinent[mapID]; ok {
		return name
	}
	return "Instances/Other"
}

var AuctionHouseNames = map[int]string{
	2: "Alliance",
	6: "Horde",
	7: "Neutral",
}

func GetAuctionHouseName(houseid int) string {
	if name, exists := AuctionHouseNames[houseid]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", houseid)
}

func GetDesertionTypeName(desertionType int) string {
	desertionTypeNames := map[int]string{
		0: "Leave",
		1: "Offline",
		2: "Desert",
		3: "Finish",
	}
	if name, exists := desertionTypeNames[desertionType]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", desertionType)
}

var WeaponSubclassNames = map[int]string{
	0: "One-Handed Axes", 1: "Two-Handed Axes", 2: "Bows", 3: "Guns",
	4: "One-Handed Maces", 5: "Two-Handed Maces", 6: "Polearms",
	7: "Two-Handed Swords", 8: "One-Handed Swords", 10: "Staves",
	13: "Fist Weapons", 15: "Daggers", 16: "Thrown", 18: "Crossbows",
	19: "Wands", 20: "Fishing Poles",
}

var InventoryTypeNames = map[int]string{
	1: "Head", 2: "Neck", 3: "Shoulder", 4: "Shirt", 5: "Chest",
	6: "Waist", 7: "Legs", 8: "Feet", 9: "Wrists", 10: "Hands",
	11: "Finger", 12: "Trinket", 13: "One-Hand", 14: "Shield",
	15: "Ranged", 16: "Back", 17: "Two-Hand", 19: "Tabard",
	20: "Robe", 21: "Main Hand", 22: "Off Hand", 23: "Held In Off-Hand",
	26: "Ranged Right", 28: "Relic",
}

var QualityNames = map[int]string{
	0: "Poor", 1: "Common", 2: "Uncommon", 3: "Rare",
	4: "Epic", 5: "Legendary", 6: "Artifact",
}

var MoneyLogTypeNames = map[int]string{
	1: "COD",
	2: "Auction_House",
	3: "Guild_Bank_Deposit",
	4: "Guild_Bank_Withdraw",
	5: "Mail",
	6: "Trade",
}

func GetMoneyLogTypeName(mtype int) string {
	if name, exists := MoneyLogTypeNames[mtype]; exists {
		return name
	}
	return fmt.Sprintf("Unknown_%d", mtype)
}

func CopperToGold(copper int64) float64 {
	return float64(copper) / 10000.0
}
