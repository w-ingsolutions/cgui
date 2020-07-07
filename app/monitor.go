package calc

import (
	"gioui.org/layout"
)

func (w *WingCal) Monitor(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana, materijalStrana, projekatStrana func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return w.UI.SaMarginom.Layout(gtx, func(gtx C) D {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(0.3, izbornikStrana),
				layout.Flexed(0.4, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx, layout.Flexed(0.5, sumaRadovaStrana), layout.Flexed(0.5, sumaMaterijalStrana))
				}),
				layout.Flexed(0.4, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Vertical,
					}.Layout(gtx, layout.Flexed(0.5, materijalStrana), layout.Flexed(0.5, projekatStrana))
				}),
			)
		})
	}
}
