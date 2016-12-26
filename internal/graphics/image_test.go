// Copyright 2016 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphics_test

import (
	"image"
	"image/color"
	"image/color/palette"
	"testing"

	. "github.com/hajimehoshi/ebiten/internal/graphics"
)

func TestCopyImage(t *testing.T) {
	cases := []struct {
		In  image.Image
		Out *image.RGBA
	}{
		{
			In: &image.Paletted{
				Pix:    []uint8{0, 1, 1, 0},
				Stride: 2,
				Rect:   image.Rect(0, 0, 2, 2),
				Palette: color.Palette([]color.Color{
					color.Transparent, color.White,
				}),
			},
			Out: &image.RGBA{
				Pix:    []uint8{0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		{
			In: &image.RGBA{
				Pix:    []uint8{0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
			Out: &image.RGBA{
				Pix:    []uint8{0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
	}
	for _, c := range cases {
		got := CopyImage(c.In)
		if got.Rect != c.Out.Rect {
			t.Errorf("Rect: %v, want: %v", got.Rect, c.Out.Rect)
		}
		size := got.Rect.Size()
		w, h := size.X, size.Y
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				got := got.At(i, j)
				want := c.Out.At(i, j)
				if got != want {
					t.Errorf("At(%d, %d): %v, want: %v", i, j, got, want)
				}
			}
		}
	}
}

func BenchmarkCopyImageRGBA(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 4096, 4096))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyImage(img)
	}
}

func BenchmarkCopyImagePaletted(b *testing.B) {
	img := image.NewPaletted(image.Rect(0, 0, 4096, 4096), palette.Plan9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyImage(img)
	}
}
