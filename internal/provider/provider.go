package provider

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New returns a *schema.Provider.
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"codename": resourceCodename(),
		},

		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return nil, fmt.Errorf("cannot seed math/rand package with Crypto RNG: %w", err)
	}

	seed := int64(binary.LittleEndian.Uint64(b[:]))

	return rand.New(rand.NewSource(seed)), nil
}
