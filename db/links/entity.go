package links

import (
	"github.com/txgruppi/site-api/shared/types"
	"gopkg.in/mgo.v2/bson"
)

type Entity struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title string        `bson:"title,omitempty" json:"title,omitempty"`
	URL   string        `bson:"url,omitempty" json:"url,omitempty"`
	Image string        `bson:"image,omitempty" json:"image,omitempty"`
	Order int           `bson:"order,omitempty" json:"order,omitempty"`
	Size  *types.Size   `bson:"size,omitempty" json:"size,omitempty"`
}
