package world

import (
	"compress/gzip"
	"os"

	"github.com/aimjel/minecraft/nbt"
)

type worldDataCreatedBy struct {
	Service, Username string
}

type worldDataMapGenerator struct {
	Software, MapGeneratorName string
}

type worldDataSpawn struct {
	X, Y, Z int16
	H, P    int8
}

type WorldData struct {
	FormatVersion                           int8
	Name                                    string
	UUID                                    []float64
	X, Y, Z                                 int16
	CreatedBy                               worldDataCreatedBy
	MapGenerator                            worldDataMapGenerator
	TimeCreated, LastAccessed, LastModified int64
	Spawn                                   worldDataSpawn
	BlockArray                              []int8
}

type World struct {
	Data WorldData
}

func (w *World) SetBlock(x1, y1, z1 int16, blockType byte) {
	x, y, z := int(x1), int(y1), int(z1)
	wx, wz := int(w.Data.X), int(w.Data.Z)

	i := x + wx*(z+wz*y)
	w.Data.BlockArray[i] = int8(blockType)
}

func LoadWorld() *World {
	d1, _ := os.Open("world/main.cw")

	dat, _ := gzip.NewReader(d1)

	var d WorldData

	nbt.NewDecoder(dat).Decode(&d)

	return &World{d}
}

func (w *World) Save() {
	file, _ := os.Create("world/main.cw")
	g := gzip.NewWriter(file)

	nbt.NewEncoder(g).Encode(w.Data)

	g.Close()
	file.Close()
}

type Block byte

const (
	BlockAir Block = iota
	BlockStone
	BlockGrass
	BlockDirt
	BlockCobblestone
	BlockPlanks
	BlockSapling
	BlockBedrock
	BlockFlowingWater
	BlockWater
	BlockFlowingLava
	BlockLava
	BlockSand
	BlockGravel
	BlockGoldOre
	BlockIronOre
	BlockCoalOre
	BlockWood
	BlockLeaves
	BlockSponge
	BlockGlass

	BlockRed
	BlockOrange
	BlockYellow
	BlockLime
	BlockGreen
	BlockTeal
	BlockAqua
	BlockCyan
	BlockBlue
	BlockIndigo
	BlockViolet
	BlockMagenta
	BlockPink
	BlockBlack
	BlockGray
	BlockWhite

	BlockDandelion
	BlockRose
	BlockBrownMushroom
	BlockRedMushroom

	BlockGold
	BlockIron
	BlockDoubleSlab
	BlockSlab
	BlockBricks
	BlockTNT
	BlockBookshelf
	BlockMoss
	BlockObsidian
)
