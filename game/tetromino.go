package game

import (
	"lets-go-tetris/render"
	"math/rand"
)

const (
	shapeX = 4
	shapeY = 4
)

// Shape 타입은 테트리스 블록의 모양 유형을 나타낸다.
type Shape int

// I		긴 막대기형
// J		ㄱ 모양
// L		L 모양 (J와 거울대칭)
// O		2x2 정사각형 블록
// S		두 번 꺾인(S) 모양
// T		T 모양
// Z		두 번 꺾인(Z) 모양 (S와 거울 대칭)
const (
	I Shape = 0 + iota
	J
	L
	O
	S
	T
	Z
)

// Rotation 타입은 미노의 회전 가능한 유형을 나타낸다.
type Rotation int

// Rotate 타입은 미노의 회전했을 때 경우의 수를 나타낸다.
type Rotate int

// Rotation 타입 총 4방향
const (
	Zero Rotation = 0 + iota
	Right
	Two
	Left
	RotationMax
)

// Rotate 타입 2^3개
const (
	ZtoR Rotate = 0 + iota
	RtoZ
	RtoT
	TtoR
	TtoL
	LtoT
	LtoZ
	ZtoL
)

const i0 = `
oooo
xxxx
oooo
oooo
`

const i1 = `
ooxo
ooxo
ooxo
ooxo
`

const i2 = `
oooo
oooo
xxxx
oooo
`

const i3 = `
oxoo
oxoo
oxoo
oxoo
`

const j0 = `
xooo
xxxo
oooo
oooo
`

const j1 = `
oxxo
oxoo
oxoo
oooo
`

const j2 = `
oooo
xxxo
ooxo
oooo
`

const j3 = `
oxoo
oxoo
xxoo
oooo
`

const s0 = `
oxxo
xxoo
oooo
oooo
`

const s1 = `
oxoo
oxxo
ooxo
oooo
`

const s2 = `
oooo
oxxo
xxoo
oooo
`

const s3 = `
xooo
xxoo
oxoo
oooo
`

const z0 = `
xxoo
oxxo
oooo
oooo
`

const z1 = `
ooxo
oxxo
oxoo
oooo
`

const z2 = `
oooo
xxoo
oxxo
oooo
`
const z3 = `
oxoo
xxoo
xooo
oooo
`

const t0 = `
oxoo
xxxo
oooo
oooo
`

const t1 = `
oxoo
oxxo
oxoo
oooo
`
const t2 = `
oooo
xxxo
oxoo
oooo
`

const t3 = `
oxoo
xxoo
oxoo
oooo
`

const o0 = `
oxxo
oxxo
oooo
oooo
`

const l0 = `
ooxo
xxxo
oooo
oooo
`

const l1 = `
oxoo
oxoo
oxxo
oooo
`

const l2 = `
oooo
xxxo
xooo
oooo
`

const l3 = `
xxoo
oxoo
oxoo
oooo
`

var shapes = [][]string{
	{i0, i1, i2, i3},
	{j0, j1, j2, j3},
	{l0, l1, l2, l3},
	{o0, o0, o0, o0},
	{s0, s1, s2, s3},
	{t0, t1, t2, t3},
	{z0, z1, z2, z3},
}

var colors = []uint32{
	0xff00d8ff,
	0xff0100ff,
	0xffffbb00,
	0xffffe400,
	0xffabf200,
	0xffff00dd,
	0xffff0000,
}

type cell bool

const (
	o cell = true
	x cell = false
)

var wallKicks = map[Rotate][][]int{
	ZtoR: {{0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}},
	RtoZ: {{0, 0}, {+1, 0}, {+1, +1}, {0, -2}, {+1, -2}},
	RtoT: {{0, 0}, {+1, 0}, {+1, +1}, {0, -2}, {+1, -2}},
	TtoR: {{0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}},
	TtoL: {{0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}},
	LtoT: {{0, 0}, {-1, 0}, {-1, +1}, {0, -2}, {-1, -2}},
	LtoZ: {{0, 0}, {-1, 0}, {-1, +1}, {0, -2}, {-1, -2}},
	ZtoL: {{0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}},
}

var iKicks = map[Rotate][][]int{
	ZtoR: {{0, 0}, {-2, 0}, {+1, 0}, {-2, +1}, {+1, -2}},
	RtoZ: {{0, 0}, {+2, 0}, {-1, 0}, {+2, -1}, {-1, +2}},
	RtoT: {{0, 0}, {-1, 0}, {+2, 0}, {-1, -2}, {+2, +1}},
	TtoR: {{0, 0}, {+1, 0}, {-2, 0}, {+1, +2}, {-2, -1}},
	TtoL: {{0, 0}, {+2, 0}, {-1, 0}, {+2, -1}, {-1, +2}},
	LtoT: {{0, 0}, {-2, 0}, {+1, 0}, {-2, +1}, {+1, -2}},
	LtoZ: {{0, 0}, {+1, 0}, {-2, 0}, {+1, +2}, {-2, -1}},
	ZtoL: {{0, 0}, {-1, 0}, {+2, 0}, {-1, -2}, {+2, +1}},
}

type tetromino struct {
	shape    Shape
	x, y     int
	cells    [][]cell
	color    uint32
	rotation Rotation
}

func randomTetromino() *tetromino {
	s := rand.Intn(len(shapes) - 1)
	m := &tetromino{color: colors[s]}
	m.init(shapes[s])
	m.shape = Shape(s)
	return m
}

func newTetromino(s Shape) *tetromino {
	m := &tetromino{color: colors[s]}
	m.init(shapes[s])
	m.shape = s
	return m
}

func (m *tetromino) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y = 0, 0
	for _, cell := range m.getCells() {
		if cell {
			infos = append(infos, &render.InfoImpl{
				PosX: m.x + x, PosY: m.y + y, Color: m.color,
			})
		}
		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}

	return infos
}

func (m *tetromino) init(rotationShapes []string) {
	m.cells = make([][]cell, RotationMax)
	for i := range m.cells {
		m.cells[i] = make([]cell, shapeX*shapeY)
	}

	for r, shape := range rotationShapes {
		i := 0
		for _, c := range shape {
			switch c {
			case 'x':
				m.cells[r][i] = true
				fallthrough
			case 'o':
				i++
			}
		}
	}
}

func (m *tetromino) rotateClockWise() Rotate {
	var r Rotate
	switch m.rotation {
	case Zero:
		r = ZtoR
	case Right:
		r = RtoT
	case Two:
		r = TtoL
	case Left:
		r = LtoZ
	}
	m.rotate(m.rotation + 1)
	return r
}

func (m *tetromino) rotateCounterClockWise() Rotate {
	var r Rotate
	switch m.rotation {
	case Zero:
		r = ZtoL
	case Right:
		r = RtoZ
	case Two:
		r = TtoR
	case Left:
		r = LtoT
	}
	m.rotate(m.rotation - 1)
	return r
}

func (m *tetromino) rotate(r Rotation) {
	m.rotation = (r%RotationMax + RotationMax) % RotationMax
}

func (m *tetromino) wallKick(g *ground, r Rotate) bool {
	if m.shape == I {
		for _, v := range iKicks[r] {
			m.x += v[0]
			m.y += v[1]
			if !g.collide(m) {
				return true
			}
			m.x -= v[0]
			m.y -= v[1]
		}
	} else {
		for _, v := range wallKicks[r] {
			m.x += v[0]
			m.y += v[1]
			if !g.collide(m) {
				return true
			}
			m.x -= v[0]
			m.y -= v[1]
		}
	}
	return false
}

func (m *tetromino) getCells() []cell {
	return m.cells[m.rotation]
}

func (m *tetromino) getPosition() (int, int) {
	return m.x, m.y
}

func (m *tetromino) getColor() uint32 {
	return m.color
}