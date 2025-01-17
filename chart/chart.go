package chart

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// GenerateImageChart kreira graf kao sliku (PNG).
func GenerateImageChart(loadTimes []float64, urls []string, output string) error {
	// Kreiraj novi plot
	p := plot.New()

	// Postavi naslov i oznake na osama
	p.Title.Text = "Performanse Učitavanja Sajtova"
	p.X.Label.Text = "Sajtovi"
	p.Y.Label.Text = "Vreme učitavanja (ms)"

	// Dodaj podatke
	pts := make(plotter.XYs, len(loadTimes))
	for i, t := range loadTimes {
		pts[i].X = float64(i)
		pts[i].Y = t
	}

	// Kreiraj linijski graf
	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Crvena boja

	// Dodaj graf na plot
	p.Add(line)

	// Dodaj oznake sajtova na X-osi
	p.NominalX(urls...)

	// Sačuvaj plot kao PNG sliku
	err = p.Save(10*vg.Inch, 5*vg.Inch, output)
	if err != nil {
		return err
	}

	return nil
}
