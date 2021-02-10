package specs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stdapps/pastebin/pastebin"
	"github.com/stretchr/testify/require"
)

// TestStorage tests concrete storage implementations
func TestStorage(t *testing.T, concrete pastebin.Storage) {
	t.Run(`Stored Paste should be retrievable`, func(t *testing.T) {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		pasteKey := fmt.Sprintf("%s%d", "key-", rand.Intn(9999))
		paste := pastebin.Paste{
			Key:  pasteKey,
			Body: "Body",
		}

		err := concrete.StorePaste(paste)
		require.Nil(t, err)

		retrievedPaste, err := concrete.GetPaste(pasteKey)
		require.Nil(t, err)

		require.Equal(t, paste, retrievedPaste)

	})
}
