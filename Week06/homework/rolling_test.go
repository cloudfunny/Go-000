package homework

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRolling(t *testing.T) {
	t.Run("for race", func(t *testing.T) {
		r := NewRolling(WithBucketDuration(2 * time.Millisecond))
		for i := 0; i < 4; i++ {
			go func() {
				for {
					time.Sleep(1 * time.Millisecond)
					r.Add(1)
				}
			}()
		}

		go func() {
			for {
				time.Sleep(1 * time.Millisecond)
				r.Sum()
			}
		}()

		time.Sleep(100 * time.Millisecond)
	})

	t.Run("long time", func(t *testing.T) {
		r := NewRolling(WithBucketDuration(2 * time.Millisecond))
		r.Add(1)
		time.Sleep(50 * time.Millisecond)
		assert.Equal(t, float64(0), r.Sum())
	})

	t.Run("short time", func(t *testing.T) {
		r := NewRolling(WithBucketDuration(10 * time.Millisecond))
		r.Add(1)
		time.Sleep(25 * time.Millisecond)
		r.Add(2)
		assert.Equal(t, float64(3), r.Sum())
		assert.Equal(t, float64(1), r.buckets[0].val)
		assert.Equal(t, float64(2), r.buckets[2].val)
	})
}
