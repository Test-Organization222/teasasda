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
	t.Run(`Paste should be retrievable after being stored`, func(t *testing.T) {
		// arrange: create a paste struct
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		pasteKey := fmt.Sprintf("%s%d", "key-", rand.Intn(9999))
		paste := pastebin.Paste{
			Key:  pasteKey,
			Body: "Body",
		}
		// store
		err := concrete.StorePaste(paste)
		require.Nil(t, err)

		// retreive
		retrievedPaste, err := concrete.GetPaste(pasteKey)
		require.Nil(t, err)

		require.Equal(t, paste, retrievedPaste)
	})

	t.Run(`Raise error if the Paste KEY already exists`, func(t *testing.T) {
		// arrange
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		pasteKey := fmt.Sprintf("%s%d", "key-", rand.Intn(9999))
		paste := pastebin.Paste{
			Key:  pasteKey,
			Body: "Body",
		}
		// store beforehand
		require.Nil(t, concrete.StorePaste(paste))
		_, err := concrete.GetPaste(pasteKey)
		require.Nil(t, err)

		// store again
		err = concrete.StorePaste(paste)

		// expect pastebin.ErrKeyExists Error
		require.ErrorIs(t, err, pastebin.ErrKeyExists)
	})
}
