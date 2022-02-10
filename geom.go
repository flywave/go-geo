package geo

import (
	"github.com/flywave/go-geom"
	"github.com/flywave/go-geom/general"

	vec2d "github.com/flywave/go3d/float64/vec2"
)

func applyCoords(coords [][]float64, src Proj, dst Proj) [][]float64 {
	pts := make([]vec2d.T, len(coords))
	for i, p := range coords {
		pts[i] = vec2d.T{p[0], p[1]}
	}
	pts = src.TransformTo(dst, pts)
	for i, p := range pts {
		coords[i][0], coords[i][1] = p[0], p[1]
	}

	return coords
}

func ApplyGeometry(g geom.Geometry, src Proj, dst Proj) geom.Geometry {
	switch t := (g).(type) {
	case geom.Point:
		pt := t.Data()
		dst := src.TransformTo(dst, []vec2d.T{{pt[0], pt[1]}})
		return general.NewPoint(dst[0][:])
	case geom.Point3:
		pt := t.Data()
		dst := src.TransformTo(dst, []vec2d.T{{pt[0], pt[1]}})
		return general.NewPoint([]float64{dst[0][0], dst[0][1], pt[2]})
	case geom.MultiPoint:
		pts := t.Data()
		pts = applyCoords(pts, src, dst)
		return general.NewMultiPoint(pts)
	case geom.MultiPoint3:
		pts := t.Data()
		pts = applyCoords(pts, src, dst)
		return general.NewMultiPoint(pts)
	case geom.LineString:
		pts := t.Data()
		pts = applyCoords(pts, src, dst)
		return general.NewLineString(pts)
	case geom.LineString3:
		pts := t.Data()
		pts = applyCoords(pts, src, dst)
		return general.NewLineString(pts)
	case geom.MultiLine:
		ls := t.Data()
		for i, pts := range ls {
			ls[i] = applyCoords(pts, src, dst)
		}
		return general.NewMultiLineString(ls)
	case geom.MultiLine3:
		ls := t.Data()
		for i, pts := range ls {
			ls[i] = applyCoords(pts, src, dst)
		}
		return general.NewMultiLineString3(ls)
	case geom.Polygon:
		ls := t.Data()
		for i, pts := range ls {
			ls[i] = applyCoords(pts, src, dst)
		}
		return general.NewPolygon(ls)
	case geom.Polygon3:
		ls := t.Data()
		for i, pts := range ls {
			ls[i] = applyCoords(pts, src, dst)
		}
		return general.NewPolygon3(ls)
	case geom.MultiPolygon:
		polys := t.Data()
		for j, ls := range polys {
			for i, pts := range ls {
				polys[j][i] = applyCoords(pts, src, dst)
			}
		}
		return general.NewMultiPolygon(polys)
	case geom.MultiPolygon3:
		polys := t.Data()
		for j, ls := range polys {
			for i, pts := range ls {
				polys[j][i] = applyCoords(pts, src, dst)
			}
		}
		return general.NewMultiPolygon3(polys)
	}
	return nil
}
