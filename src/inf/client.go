package inf

import (
	"os"

	wishlistlib "github.com/Sam36502/WishlistLib-go"
)

var g_WishClient *wishlistlib.WishClient = nil

func InitWishlistClient() {
	baseURL, found := os.LookupEnv("WISHLIST_BASE_URL")
	if !found {
		baseURL = "http://localhost"
	}

	g_WishClient = &wishlistlib.WishClient{
		BaseURL: baseURL,
		Port:    wishlistlib.DEFAULT_PORT,
		Token:   wishlistlib.Token{},
	}
}

func GetWishlistClient(tok wishlistlib.Token) *wishlistlib.WishClient {
	if g_WishClient == nil {
		InitWishlistClient()
	}

	return g_WishClient
}
