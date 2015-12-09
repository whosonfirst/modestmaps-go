package modestmaps

func NewTemplatedMercatorProvider (template string) *TemplatedMercatorProvider {

     transformation = DeriveTransformation(-math.PI, math.PI, 0.0, 0.0, math.PI, math.PI, 1.0, 0.0, -math.PI, -math.PI, 0.0, 1.0)
     projection := NewProjectionWithTransformation(0, tranformation)

     return &TemplatedMercatorProvider{projection, template}
}

type TemplatedMercatorProvider struct {
     projection *IProjection
     template string
}

func (t *TemplatedMercatorProjection) TileHeight int {
     return 256
}

func (t *TemplatedMercatorProjection) TileWidth int {
     return 256
}

func (t *TemplatedMercatorProjection) GetTileUrls (c *Coordinate) []string {
     urls := make([]string, 0)

     // please write me

     return urls
}
