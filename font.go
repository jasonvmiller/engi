// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engi

type glyph struct {
	region   *Region
	xoffset  float32
	yoffset  float32
	xadvance float32
}

type Font struct {
	Scale    *Point
	cellHeight int
	cellWidth int
	glyphs map[rune]*glyph
}

func NewGridFont(texture *Texture, cellWidth, cellHeight int) *Font {
	i := 0
	glyphs := make(map[rune]*glyph)

	for y := 0; y < int(texture.Height())/cellHeight; y++ {
		for x := 0; x < int(texture.Width())/cellWidth; x++ {
			g := &glyph{xadvance: float32(cellWidth)}
			g.region = NewRegion(texture, x*cellWidth, y*cellHeight, cellWidth, cellHeight)
			glyphs[rune(i)] = g
			i += 1
		}
	}

	return &Font{glyphs: glyphs, cellHeight: cellHeight, cellWidth: cellWidth}
}

func (f *Font) Remap(mapping string) {
	glyphs := make(map[rune]*glyph)

	i := 0
	for _, v := range mapping {
		glyphs[v] = f.glyphs[rune(i)]
		i++
	}

	f.glyphs = glyphs
}

func (f *Font) Put(batch *Batch, r rune, x, y float32, color uint32) {
	scaleX := float32(1)
	scaleY := float32(1)
	if f.Scale != nil {
		scaleX = f.Scale.X
		scaleY = f.Scale.Y
	}
	if g, ok := f.glyphs[r]; ok {
		batch.Draw(g.region, x+g.xoffset, y+g.yoffset, 0, 0, scaleX, scaleY, 0, color, 1)
	}
}

func (f *Font) Print(batch *Batch, text string, x, y float32, color uint32) {
	xx := x
	scaleX := float32(1)
	scaleY := float32(1)
	if f.Scale != nil {
		scaleX = f.Scale.X
		scaleY = f.Scale.Y
	}
	for _, r := range text {
		if g, ok := f.glyphs[r]; ok {
			batch.Draw(g.region, xx+g.xoffset, y+g.yoffset, 0, 0, scaleX, scaleY, 0, color, 1)
			xx += g.xadvance * scaleX
		}
	}
}

func (f *Font) CellHeight() ( int ) {
	return int(float32(f.cellHeight) * f.Scale.Y)
}

func (f *Font) CellWidth() ( int ) {
	return int(float32(f.cellWidth) * f.Scale.X)
}
