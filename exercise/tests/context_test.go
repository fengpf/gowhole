package tests

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"gowhole/src/go1.10.1/context"
)

func Test_cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	f := func(c context.Context) {
		for i := 0; i < 100; i++ {
			if i == 50 {
				<-c.Done()
			}
			fmt.Println(i)
		}
	}
	f(ctx)
}

func Test_withCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func Test_withDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	// Output:
	// context deadline exceeded
}

func Test_withTimeout(t *testing.T) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	// Output:
	// context deadline exceeded
}

func Test_withValue(t *testing.T) {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	// Output:
	// found value: Go
	// key not found: color
}

// otherContext is a Context that's not one of the types defined in context.go.
// This lets us test code paths that differ based on the underlying type of the
// Context.
type otherContext struct {
	context.Context
}

func TestBackground(t *testing.T) {
	c := context.Background()
	if c == nil {
		t.Fatalf("Background returned nil")
	}
	select {
	case x := <-c.Done():
		t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}
	if got, want := fmt.Sprint(c), "context.Background"; got != want {
		t.Errorf("Background().String() = %q want %q", got, want)
	}
}

func TestTODO(t *testing.T) {
	c := context.TODO()
	if c == nil {
		t.Fatalf("TODO returned nil")
	}
	select {
	case x := <-c.Done():
		t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}
	if got, want := fmt.Sprint(c), "context.TODO"; got != want {
		t.Errorf("TODO().String() = %q want %q", got, want)
	}
}

func TestWithCancel(t *testing.T) {
	c1, cancel := context.WithCancel(context.Background())

	if got, want := fmt.Sprint(c1), "context.Background.WithCancel"; got != want {
		t.Errorf("c1.String() = %q want %q", got, want)
	}

	o := otherContext{c1}
	c2, _ := context.WithCancel(o)
	contexts := []context.Context{c1, o, c2}

	for i, c := range contexts {
		if d := c.Done(); d == nil {
			t.Errorf("c[%d].Done() == %v want non-nil", i, d)
		}
		if e := c.Err(); e != nil {
			t.Errorf("c[%d].Err() == %v want nil", i, e)
		}

		select {
		case x := <-c.Done():
			t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
		default:
		}
	}

	cancel()
	time.Sleep(100 * time.Millisecond) // let cancelation propagate

	for i, c := range contexts {
		select {
		case <-c.Done():
		default:
			t.Errorf("<-c[%d].Done() blocked, but shouldn't have", i)
		}
		if e := c.Err(); e != context.Canceled {
			t.Errorf("c[%d].Err() == %v want %v", i, e, context.Canceled)
		}
	}
}

func testDeadline(c context.Context, name string, failAfter time.Duration, t *testing.T) {
	select {
	case <-time.After(failAfter):
		t.Fatalf("%s: context should have timed out", name)
	case <-c.Done():
	}
	if e := c.Err(); e != context.DeadlineExceeded {
		t.Errorf("%s: c.Err() == %v; want %v", name, e, context.DeadlineExceeded)
	}
}

func TestDeadline(t *testing.T) {
	c, _ := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
	if got, prefix := fmt.Sprint(c), "context.Background.WithDeadline("; !strings.HasPrefix(got, prefix) {
		t.Errorf("c.String() = %q want prefix %q", got, prefix)
	}
	testDeadline(c, "WithDeadline", time.Second, t)

	c, _ = context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
	o := otherContext{c}
	testDeadline(o, "WithDeadline+otherContext", time.Second, t)

	c, _ = context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
	o = otherContext{c}
	c, _ = context.WithDeadline(o, time.Now().Add(4*time.Second))
	testDeadline(c, "WithDeadline+otherContext+WithDeadline", 2*time.Second, t)

	c, _ = context.WithDeadline(context.Background(), time.Now().Add(-time.Millisecond))
	testDeadline(c, "WithDeadline+inthepast", time.Second, t)

	c, _ = context.WithDeadline(context.Background(), time.Now())
	testDeadline(c, "WithDeadline+now", time.Second, t)
}

func TestTimeout(t *testing.T) {
	c, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	if got, prefix := fmt.Sprint(c), "context.Background.WithDeadline("; !strings.HasPrefix(got, prefix) {
		t.Errorf("c.String() = %q want prefix %q", got, prefix)
	}
	testDeadline(c, "WithTimeout", time.Second, t)

	c, _ = context.WithTimeout(context.Background(), 50*time.Millisecond)
	o := otherContext{c}
	testDeadline(o, "WithTimeout+otherContext", time.Second, t)

	c, _ = context.WithTimeout(context.Background(), 50*time.Millisecond)
	o = otherContext{c}
	c, _ = context.WithTimeout(o, 3*time.Second)
	testDeadline(c, "WithTimeout+otherContext+WithTimeout", 2*time.Second, t)
}

func TestCanceledTimeout(t *testing.T) {
	c, _ := context.WithTimeout(context.Background(), time.Second)
	o := otherContext{c}
	c, cancel := context.WithTimeout(o, 2*time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond) // let cancelation propagate
	select {
	case <-c.Done():
	default:
		t.Errorf("<-c.Done() blocked, but shouldn't have")
	}
	if e := c.Err(); e != context.Canceled {
		t.Errorf("c.Err() == %v want %v", e, context.Canceled)
	}
}

type key1 int
type key2 int

var k1 = key1(1)
var k2 = key2(1) // same int as k1, different type
var k3 = key2(3) // same type as k2, different int

func TestValues(t *testing.T) {
	check := func(c context.Context, nm, v1, v2, v3 string) {
		if v, ok := c.Value(k1).(string); ok == (len(v1) == 0) || v != v1 {
			t.Errorf(`%s.Value(k1).(string) = %q, %t want %q, %t`, nm, v, ok, v1, len(v1) != 0)
		}
		if v, ok := c.Value(k2).(string); ok == (len(v2) == 0) || v != v2 {
			t.Errorf(`%s.Value(k2).(string) = %q, %t want %q, %t`, nm, v, ok, v2, len(v2) != 0)
		}
		if v, ok := c.Value(k3).(string); ok == (len(v3) == 0) || v != v3 {
			t.Errorf(`%s.Value(k3).(string) = %q, %t want %q, %t`, nm, v, ok, v3, len(v3) != 0)
		}
	}

	c0 := context.Background()
	check(c0, "c0", "", "", "")

	c1 := context.WithValue(context.Background(), k1, "c1k1")
	check(c1, "c1", "c1k1", "", "")

	if got, want := fmt.Sprint(c1), `context.Background.WithValue(1, "c1k1")`; got != want {
		t.Errorf("c.String() = %q want %q", got, want)
	}

	c2 := context.WithValue(c1, k2, "c2k2")
	check(c2, "c2", "c1k1", "c2k2", "")

	c3 := context.WithValue(c2, k3, "c3k3")
	check(c3, "c2", "c1k1", "c2k2", "c3k3")

	c4 := context.WithValue(c3, k1, nil)
	check(c4, "c4", "", "c2k2", "c3k3")

	o0 := otherContext{context.Background()}
	check(o0, "o0", "", "", "")

	o1 := otherContext{context.WithValue(context.Background(), k1, "c1k1")}
	check(o1, "o1", "c1k1", "", "")

	o2 := context.WithValue(o1, k2, "o2k2")
	check(o2, "o2", "c1k1", "o2k2", "")

	o3 := otherContext{c4}
	check(o3, "o3", "", "c2k2", "c3k3")

	o4 := context.WithValue(o3, k3, nil)
	check(o4, "o4", "", "c2k2", "")
}

func XTestAllocs(t *testing.T, testingShort func() bool, testingAllocsPerRun func(int, func()) float64) {
	bg := context.Background()
	for _, test := range []struct {
		desc       string
		f          func()
		limit      float64
		gccgoLimit float64
	}{
		{
			desc:       "Background()",
			f:          func() { context.Background() },
			limit:      0,
			gccgoLimit: 0,
		},
		{
			desc: fmt.Sprintf("WithValue(bg, %v, nil)", k1),
			f: func() {
				c := context.WithValue(bg, k1, nil)
				c.Value(k1)
			},
			limit:      3,
			gccgoLimit: 3,
		},
		{
			desc: "WithTimeout(bg, 15*time.Millisecond)",
			f: func() {
				c, _ := context.WithTimeout(bg, 15*time.Millisecond)
				<-c.Done()
			},
			limit:      8,
			gccgoLimit: 15,
		},
		{
			desc: "WithCancel(bg)",
			f: func() {
				c, cancel := context.WithCancel(bg)
				cancel()
				<-c.Done()
			},
			limit:      5,
			gccgoLimit: 8,
		},
		{
			desc: "WithTimeout(bg, 5*time.Millisecond)",
			f: func() {
				c, cancel := context.WithTimeout(bg, 5*time.Millisecond)
				cancel()
				<-c.Done()
			},
			limit:      8,
			gccgoLimit: 25,
		},
	} {
		limit := test.limit
		if runtime.Compiler == "gccgo" {
			// gccgo does not yet do escape analysis.
			// TOOD(iant): Remove this when gccgo does do escape analysis.
			limit = test.gccgoLimit
		}
		numRuns := 100
		if testingShort() {
			numRuns = 10
		}
		if n := testingAllocsPerRun(numRuns, test.f); n > limit {
			t.Errorf("%s allocs = %f want %d", test.desc, n, int(limit))
		}
	}
}

func XTestSimultaneousCancels(t *testing.T) {
	root, cancel := context.WithCancel(context.Background())
	m := map[context.Context]context.CancelFunc{root: cancel}
	q := []context.Context{root}
	// Create a tree of contexts.
	for len(q) != 0 && len(m) < 100 {
		parent := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			ctx, cancel := context.WithCancel(parent)
			m[ctx] = cancel
			q = append(q, ctx)
		}
	}
	// Start all the cancels in a random order.
	var wg sync.WaitGroup
	wg.Add(len(m))
	for _, cancel := range m {
		go func(cancel context.CancelFunc) {
			cancel()
			wg.Done()
		}(cancel)
	}
	// Wait on all the contexts in a random order.
	for ctx := range m {
		select {
		case <-ctx.Done():
		case <-time.After(1 * time.Second):
			buf := make([]byte, 10<<10)
			n := runtime.Stack(buf, true)
			t.Fatalf("timed out waiting for <-ctx.Done(); stacks:\n%s", buf[:n])
		}
	}
	// Wait for all the cancel functions to return.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		buf := make([]byte, 10<<10)
		n := runtime.Stack(buf, true)
		t.Fatalf("timed out waiting for cancel functions; stacks:\n%s", buf[:n])
	}
}

func TestInterlockedCancels(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	go func() {
		parent.Done()
		cancelChild()
	}()
	cancelParent()
	select {
	case <-child.Done():
	case <-time.After(1 * time.Second):
		buf := make([]byte, 10<<10)
		n := runtime.Stack(buf, true)
		t.Fatalf("timed out waiting for child.Done(); stacks:\n%s", buf[:n])
	}
}

func TestLayersCancel(t *testing.T) {
	testLayers(t, time.Now().UnixNano(), false)
}

func TestLayersTimeout(t *testing.T) {
	testLayers(t, time.Now().UnixNano(), true)
}

func testLayers(t *testing.T, seed int64, testTimeout bool) {
	rand.Seed(seed)
	errorf := func(format string, a ...interface{}) {
		t.Errorf(fmt.Sprintf("seed=%d: %s", seed, format), a...)
	}
	const (
		timeout   = 200 * time.Millisecond
		minLayers = 30
	)
	type value int
	var (
		vals      []*value
		cancels   []context.CancelFunc
		numTimers int
		ctx       = context.Background()
	)
	for i := 0; i < minLayers || numTimers == 0 || len(cancels) == 0 || len(vals) == 0; i++ {
		switch rand.Intn(3) {
		case 0:
			v := new(value)
			ctx = context.WithValue(ctx, v, v)
			vals = append(vals, v)
		case 1:
			var cancel context.CancelFunc
			ctx, cancel = context.WithCancel(ctx)
			cancels = append(cancels, cancel)
		case 2:
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			cancels = append(cancels, cancel)
			numTimers++
		}
	}
	checkValues := func(when string) {
		for _, key := range vals {
			if val := ctx.Value(key).(*value); key != val {
				errorf("%s: ctx.Value(%p) = %p want %p", when, key, val, key)
			}
		}
	}
	select {
	case <-ctx.Done():
		errorf("ctx should not be canceled yet")
	default:
	}
	if s, prefix := fmt.Sprint(ctx), "context.Background."; !strings.HasPrefix(s, prefix) {
		t.Errorf("ctx.String() = %q want prefix %q", s, prefix)
	}
	t.Log(ctx)
	checkValues("before cancel")
	if testTimeout {
		select {
		case <-ctx.Done():
		case <-time.After(timeout + time.Second):
			errorf("ctx should have timed out")
		}
		checkValues("after timeout")
	} else {
		cancel := cancels[rand.Intn(len(cancels))]
		cancel()
		select {
		case <-ctx.Done():
		default:
			errorf("ctx should be canceled")
		}
		checkValues("after cancel")
	}
}

func TestWithCancelCanceledParent(t *testing.T) {
	parent, pcancel := context.WithCancel(context.Background())
	pcancel()
	c, _ := context.WithCancel(parent)
	select {
	case <-c.Done():
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for Done")
	}
	if got, want := c.Err(), context.Canceled; got != want {
		t.Errorf("child not cancelled; got = %v, want = %v", got, want)
	}
}

func TestWithValueChecksKey(t *testing.T) {
	panicVal := recoveredValue(func() { context.WithValue(context.Background(), []byte("foo"), "bar") })
	if panicVal == nil {
		t.Error("expected panic")
	}
	panicVal = recoveredValue(func() { context.WithValue(context.Background(), nil, "bar") })
	if got, want := fmt.Sprint(panicVal), "nil key"; got != want {
		t.Errorf("panic = %q; want %q", got, want)
	}
}

func recoveredValue(fn func()) (v interface{}) {
	defer func() { v = recover() }()
	fn()
	return
}

func TestDeadlineExceededSupportsTimeout(t *testing.T) {
	i, ok := context.DeadlineExceeded.(interface {
		Timeout() bool
	})
	if !ok {
		t.Fatal("DeadlineExceeded does not support Timeout interface")
	}
	if !i.Timeout() {
		t.Fatal("wrong value for timeout")
	}
}
