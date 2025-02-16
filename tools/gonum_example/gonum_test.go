package gonum_example

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"testing"
)

func TestCorrelation(t *testing.T) {
	x := []float64{1.9, 2.0, 5.6, 5.2, 2.3}
	y := []float64{2.0, 2.5, 19, 17, 3}
	weights := []float64{1, 1, 1, 1, 1} // 可选权重

	// 计算相关系数
	r := stat.Correlation(x, y, weights)
	fmt.Printf("皮尔逊相关系数: %v\n", r)
}

func TestCorrelationDemo1(t *testing.T) {
	x := []float64{1.9, 1.8, 1.85, 1.9, 2.0, 5.6, 5.2, 2.3}
	y := []float64{2.1, 2.0, 1.9, 2.0, 2.5, 19, 17, 3}

	// 计算相关系数
	r := stat.Correlation(x, y, nil)
	fmt.Printf("皮尔逊相关系数: %v\n", r)
}

func TestCorrelationDemo(t *testing.T) {
	x := []float64{1.9, 2.0, 5.6, 10, 11}
	y := []float64{2.0, 2.5, 8, 13, 15}
	weights := []float64{1, 1, 1, 1, 1} // 可选权重

	// 计算相关系数
	r := stat.Correlation(x, y, weights)
	fmt.Printf("皮尔逊相关系数: %v\n", r)

	z := []float64{1.0, 1.1, 1.01, 1.3, 1.3}
	// 计算相关系数
	r = stat.Correlation(x, z, weights)
	fmt.Printf("皮尔逊相关系数: %v\n", r)

}

func TestPlotDemo(t *testing.T) {
	// 创建图表
	p := plot.New()
	p.Title.Text = "Scatter Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 添加数据点
	points := plotter.XYs{
		{X: 1, Y: 1},
		{X: 2, Y: 4},
		{X: 3, Y: 9},
		{X: 4, Y: 16},
	}

	scatter, _ := plotter.NewScatter(points)
	p.Add(scatter)

	// 保存图表
	p.Save(100*vg.Inch, 100*vg.Inch, "scatter.png")
}
