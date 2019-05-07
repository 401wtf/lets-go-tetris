package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/render"
)

var _ = Describe("ground reset 테스트", func() {
	type testData struct {
		x, y     int
		expected []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		g := ground{width: d.x, height: d.y}
		g.reset()
		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("2x3", testData{2, 3, []cell{
			x, x,
			x, x,
			x, x,
		}}),
		Entry("4x3", testData{4, 3, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
		}}),
	)
})

var _ = Describe("ground merge 테스트", func() {
	type testData struct {
		s        Shape
		x, y     int
		expected []cell
	}

	g := ground{width: 4, height: 10}
	g.reset()

	DescribeTable("테스트 케이스", func(d testData) {
		m := tetromino{x: d.x, y: d.y}
		m.init(shapes[d.s])
		g.merge(&m)

		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("L", testData{L, 0, 0, []cell{
			x, x, o, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("I", testData{I, 2, 3, []cell{
			x, x, o, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
			x, x, o, o,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("O", testData{O, 3, 5, []cell{
			x, x, o, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
			x, x, o, o,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
		}}),
	)
})

var _ = Describe("ground collide 테스트", func() {
	It("정상적인 상황에서는 충돌하지 않는다.", func() {
		g := ground{width: 4, height: 10}
		g.reset()

		m := newTetromino(S)
		g.merge(m)

		m = newTetromino(L)
		m.y = 4

		actual := g.collide(m)

		Expect(actual).Should(Equal(false))
	})

	It("이미 ground에 머지되어있는 블럭과 잘 충돌한다.", func() {
		g := ground{width: 4, height: 10}
		g.reset()

		m := newTetromino(S)
		g.merge(m)

		m = newTetromino(L)
		actual := g.collide(m)

		Expect(actual).Should(Equal(true))
	})

	It("ground 밖으로 나갔는지 체크한다.", func() {
		g := ground{width: 4, height: 10}
		g.reset()

		m := newTetromino(I)
		m.x = 100
		m.y = 100

		actual := g.collide(m)

		Expect(actual).Should(Equal(true))
	})
})

/*
var _ = Describe("ground step 테스트", func() {
	g := ground{width: 4, height: 10}
	g.reset()

	m := newTetromino(I)
	m.y = 6
	g.merge(m)

	m = newTetromino(Z)
	m.y = 4

	It("블럭이 충돌하지 않으면서 머지되지 않고 한칸 내려갔다.", func() {
		Expect(g.step(m)).Should(Equal(false))

		expected := []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}
		var y int
		y = 5

		actual := g.cells
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
		Expect(m.y).Should(Equal(y))
	})

	It("블럭이 충돌하면서 머지된다.", func() {
		Expect(g.step(m)).Should(Equal(true))

		expected := []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, x, x,
			x, o, o, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}

		actual := g.cells
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
	})
})
*/

var _ = Describe("ground removeLines 테스트", func() {
	type testData struct {
		expected int
		before   []cell
		after    []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		g := ground{width: 4, height: 10}
		g.reset()
		g.cells = d.before

		actual := g.removeLines()
		Expect(actual).Should(Equal(d.expected))

		diff := deep.Equal(g.cells, d.after)
		Expect(diff).Should(BeNil())
	},
		Entry("1줄", testData{1, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, o, x, o,
			o, o, o, o,
			o, o, x, o,
		}, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, o, x, o,
			o, o, x, o,
		}}),
		Entry("3줄", testData{3, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, o, o,
			o, o, o, o,
			o, o, x, o,
			o, o, o, o,
		}, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, x, o,
		}}),
		Entry("4줄", testData{4, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, o, x, x,
			x, o, x, x,
			o, o, o, o,
			o, o, o, o,
			o, o, o, o,
			o, o, o, o,
		}, []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, o, x, x,
			x, o, x, x,
		}}),
	)
})

var _ = Describe("ground.RenderInfo() 함수가", func() {
	It("렌더링 정보를 제대로 반환한다.", func() {
		g := ground{
			width:  2,
			height: 3,
			cells:  []cell{o, o, x, o, o, o},
			colors: []uint32{123, 123, 123, 123, 123, 123},
		}

		expected := []render.Info{
			&render.InfoImpl{PosX: 0, PosY: 0, Color: 123},
			&render.InfoImpl{PosX: 1, PosY: 0, Color: 123},
			&render.InfoImpl{PosX: 0, PosY: 1, Color: 0xff353535},
			&render.InfoImpl{PosX: 1, PosY: 1, Color: 123},
			&render.InfoImpl{PosX: 0, PosY: 2, Color: 123},
			&render.InfoImpl{PosX: 1, PosY: 2, Color: 123},
		}
		actual := g.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})