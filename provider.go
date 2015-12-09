package modestmaps

type IMapProvider interface {
	GetTileUrls(c *Coordinate)
	TileWidth() int
	TileHeight() int
	LocationCoordinate(l *Location) *Coordinate
	CoordinateLocation(c *Coorindate) *Location
	SourceCoordinate(c *Coordinate)
}
