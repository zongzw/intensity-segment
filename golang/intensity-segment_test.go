package main

import (
	"reflect"
	"testing"
)

// TestIntensitySegments_set Test IntensitySegments.set
//
//	Notice the comments:
//		____ means existing IntensitySegments
//		---- means new set segment
//		==== means overlap between ____ and ----
//
// In Chinese:
//
// 这是测试IntensitySegments set 的模版函数：
//
//	____ 表示 已有的IntensitySegments
//	---- 表示 调用新的set函数设置segment
//	==== 表示 新设置的segment与原有segment有重叠
func TestIntensitySegments_set(t *testing.T) {
	IntensitySegments_set_tmpl(t, 10, 50, 1, 5, 8, 2, "[[5 2] [8 0] [10 1] [50 0]]")     // ---- ____
	IntensitySegments_set_tmpl(t, 10, 50, 1, 5, 20, 2, "[[5 2] [20 1] [50 0]]")          // --=====__
	IntensitySegments_set_tmpl(t, 10, 50, 1, 5, 60, 2, "[[5 2] [60 0]]")                 // --=====--
	IntensitySegments_set_tmpl(t, 10, 50, 1, 20, 30, 2, "[[10 1] [20 2] [30 1] [50 0]]") // __=====__
	IntensitySegments_set_tmpl(t, 10, 50, 1, 20, 60, 2, "[[10 1] [20 2] [60 0]]")        // __=====--
	IntensitySegments_set_tmpl(t, 10, 50, 1, 55, 60, 2, "[[10 1] [50 0] [55 2] [60 0]]") // ____ ----
	IntensitySegments_set_tmpl(t, 10, 50, 1, 10, 30, 2, "[[10 2] [30 1] [50 0]]")        // ====_____
}

// TestIntensitySegments_add Test IntensitySegments.add
//
//	Notice the comments:
//		____ means existing IntensitySegments
//		---- means new add segment
//		==== means overlap between ____ and ----
//
// In Chinese:
//
// 这是测试IntensitySegments add 的模版函数：
//
//	____ 表示 已有的IntensitySegments
//	---- 表示 调用新的add函数设置segment
//	==== 表示 新设置的segment与原有segment有重叠
func TestIntensitySegments_add(t *testing.T) {
	IntensitySegments_add_tmpl(t, 10, 50, 1, 5, 8, 2, "[[5 2] [8 0] [10 1] [50 0]]")     // ---- ____
	IntensitySegments_add_tmpl(t, 10, 50, 1, 5, 20, 2, "[[5 2] [10 3] [20 1] [50 0]]")   // --=====__
	IntensitySegments_add_tmpl(t, 10, 50, 1, 5, 60, 2, "[[5 2] [10 3] [50 2] [60 0]]")   // --=====--
	IntensitySegments_add_tmpl(t, 10, 50, 1, 20, 30, 2, "[[10 1] [20 3] [30 1] [50 0]]") // __=====__
	IntensitySegments_add_tmpl(t, 10, 50, 1, 20, 60, 2, "[[10 1] [20 3] [50 2] [60 0]]") // __=====--
	IntensitySegments_add_tmpl(t, 10, 50, 1, 55, 60, 2, "[[10 1] [50 0] [55 2] [60 0]]") // ____ ----
	IntensitySegments_add_tmpl(t, 10, 50, 1, 10, 30, 2, "[[10 3] [30 1] [50 0]]")        // ====_____
}

// IntensitySegments_set_tmpl Unit Test template function
// 比较两次set之后是否符合预期
// is.set(f1, t1, a1)
// is.set(f2, t2, a2)
func IntensitySegments_set_tmpl(t *testing.T, f1, t1, a1, f2, t2, a2 int, expected string) {
	is := NewIntensitySegments()
	is.set(f1, t1, a1)
	is.set(f2, t2, a2)
	a := is.dumps()
	b := expected
	if !reflect.DeepEqual(a, b) {
		t.Errorf("assert failure: a: %v, b: %v", a, b)
	}
}

// IntensitySegments_add_tmpl Unit Test template function
// 比较一次add之后是否符合预期
// is.set(f1, t1, a1)
// is.add(f2, t2, a2)
func IntensitySegments_add_tmpl(t *testing.T, f1, t1, a1, f2, t2, a2 int, expected string) {
	is := NewIntensitySegments()
	is.set(f1, t1, a1)
	is.add(f2, t2, a2)
	a := is.dumps()
	b := expected
	if !reflect.DeepEqual(a, b) {
		t.Errorf("assert failure: a: %v, b: %v", a, b)
	}
}
