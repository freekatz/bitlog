package types

import (
	"sort"
	"time"
)

// TODO 精简，目前看来还用不上 timeline 的

type (
	Timestamp int64
	TimePoint struct {
		Timestamp
		Log
	}
	Timeline struct {
		TimePointSet []TimePoint
		Since        Timestamp
	}
)

func FromTime(t time.Time) Timestamp {
	return Timestamp(t.UnixNano())
}

func FromNow() Timestamp {
	return FromTime(time.Now())
}

func FromInt64(t int64) Timestamp {
	return Timestamp(t)
}

// TODO add more method about Timestamp
func (t Timestamp) FormatString() string {
	return ""
}

// TODO add more method about Timestamp
func (t Timestamp) FormatTime() time.Time {
	return time.Now()
}

func NewTimeline(point TimePoint) Timeline {
	tl := Timeline{make([]TimePoint, 0), FromInt64(0)}
	tl.TimePointSet = append(tl.TimePointSet, point)
	tl.Since = point.Timestamp
	return tl
}

// PushPoint push a point into point set
func (tl *Timeline) PushPoint(points ...TimePoint) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Timestamp < points[j].Timestamp
	})
	if len(tl.TimePointSet) == 0 {
		tl.Since = points[0].Timestamp
	}
	tl.TimePointSet = append(tl.TimePointSet, points...)
}

// GetPointByRange find range TimestampStatus between start and end
func (tl *Timeline) GetPointByRange(start, end Timestamp) Timeline {
	n := len(tl.TimePointSet)
	l := sort.Search(n, func(i int) bool {
		return tl.TimePointSet[i].Timestamp <= start
	})
	r := sort.Search(n, func(i int) bool {
		return tl.TimePointSet[i].Timestamp >= end
	})
	if l >= n || l > r {
		return Timeline{}
	}
	return Timeline{tl.TimePointSet[l : r+1], tl.TimePointSet[l].Timestamp}
}

// GetPointByTime find left and right TimestampStatus about the pivot
func (tl *Timeline) GetPointByTime(pivot Timestamp) Timeline {
	n := len(tl.TimePointSet)
	r := sort.Search(n, func(i int) bool {
		return tl.TimePointSet[i].Timestamp >= pivot
	})
	if r >= n || r <= 0 {
		return Timeline{}
	}
	return Timeline{tl.TimePointSet[r-1 : r+1], tl.TimePointSet[r-1].Timestamp}
}

// GetPointByType find the TimestampStatus that Type in ts
func (tl *Timeline) GetPointByType(t RawLogType) Timeline {
	_tpSet := []TimePoint{}
	for i := range tl.TimePointSet {
		lo := tl.TimePointSet[i].Log
		if lo.Raw().Type == t {
			_tpSet = append(_tpSet, tl.TimePointSet[i])
		}
	}
	if len(_tpSet) == 0 {
		return Timeline{}
	}
	return Timeline{_tpSet, _tpSet[0].Timestamp}
}

// GetPointSince find the TimestampStatus Since the time point
func (tl *Timeline) GetPointSince(pivot Timestamp) Timeline {
	n := len(tl.TimePointSet)
	r := sort.Search(n, func(i int) bool {
		return tl.TimePointSet[i].Timestamp >= pivot
	})
	if r >= n || r <= 0 {
		return Timeline{}
	}
	return Timeline{tl.TimePointSet[r:], tl.TimePointSet[r].Timestamp}
}

// GetPointBefore find the TimestampStatus Before the time point
func (tl *Timeline) GetPointBefore(pivot Timestamp) Timeline {
	n := len(tl.TimePointSet)
	r := sort.Search(n, func(i int) bool {
		return tl.TimePointSet[i].Timestamp >= pivot
	})
	if r >= n || r <= 0 {
		return Timeline{}
	}
	return Timeline{tl.TimePointSet[:r], tl.Since}
}
