package main

import (
	"sync"
	"time"
)

//令牌桶算法的基本过程如下：
//
//　　假如用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中；
//
//　　假设桶最多可以存发b个令牌。如果令牌到达时令牌桶已经满了，那么这个令牌会被丢弃；
//
//　　当一个n个字节的数据包到达时，就从令牌桶中删除n个令牌，并且数据包被发送到网络；
//
//　　如果令牌桶中少于n个令牌，那么不会删除令牌，并且认为这个数据包在流量限制之外；
//
//　　算法允许最长b个字节的突发，但从长期运行结果看，数据包的速率被限制成常量r。对于在流量限制外的数据包可以以不同的方式处理：
//
//　　它们可以被丢弃；
//
//　　它们可以排放在队列中以便当令牌桶中累积了足够多的令牌时再传输；
//
//　　它们可以继续发送，但需要做特殊标记，网络过载的时候将这些特殊标记的包丢弃。

func main() {
}

type Limiter struct {
	rate   float64 //生成速率
	bucket int     //桶的大小

	tokens    float64   //桶中目前剩余的token数目，可以为负
	last      time.Time //上一次限流器更新tokens字段的时间
	lastEvent time.Time //最近一次限流事件的时间(过去或者将来)

	mu sync.Mutex
}

func Rate(interval time.Duration) float64 {
	return 1 / interval.Seconds()
}

func NewLimiter(r float64, b int) *Limiter {
	return &Limiter{
		rate:   r,
		bucket: b,
	}
}

//拿令牌
//1、计算从上次取 Token 的时间到当前时刻，期间一共新产生了多少 Token：
//我们只在取 Token 之前生成新的 Token，也就意味着每次取 Token 的间隔，实际上也是生成 Token 的间隔
//当前 Token 数目 = 新产生的 Token 数目 + 之前剩余的 Token 数目 - 要消费的 Token 数目

//2、如果消费后剩余 Token 数目大于零，说明此时 Token 桶内仍不为空，此时 Token 充足，无需调用侧等待。
//如果 Token 数目小于零，则需等待一段时间。
//那么这个时候，我们可以利用 durationFromTokens 将当前负值的 Token 数转化为需要等待的时间。
//
//3、将需要等待的时间等相关结果返回给调用方。

// @param n 要消费的token数量
// @param maxWait 愿意等待的最长时间
func (l *Limiter) Take(now time.Time, n int, maxWaitDuration time.Duration) bool {
	l.mu.Lock()

	now, last, tokens := l.advance(now)

	//看下取完之后，桶还能剩能下多少token
	tokens -= float64(n)

	var waitDuration time.Duration
	if tokens < 0 { //如token数目为负，则说明token不够，需要等待一段时间
		waitDuration = l.durationFromTokens(-tokens)
	}

	ok := n <= l.bucket && waitDuration <= maxWaitDuration

	if ok { //更新桶的token数目，更新last，lastEvent 时间，
		l.last = now
		l.tokens = tokens
		l.lastEvent = now.Add(waitDuration)
	} else {
		l.last = last
	}

	l.mu.Unlock()
	return ok
}

// @param now
// @return newNow 似乎还是这个now，没变
// @return newLast 如果 last > now, 则last为now
// @return newTokens 当前桶中应有的数目
func (l *Limiter) advance(now time.Time) (newNow time.Time, newLast time.Time, newTokens float64) {
	//last 代表上一个取token的时间
	last := l.last
	if now.Before(last) {
		last = now
	}

	//算把桶填满的时间 maxElapsed
	maxElapsed := l.durationFromTokens(float64(l.bucket) - l.tokens)

	//表示从当前到上次取token过去多长时间，并且不能大于将桶填满的时间
	elapsed := now.Sub(last)
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	delta := l.tokensFromDuration(elapsed) //过去这段时间总共产生了多少token

	tokens := l.tokens + delta
	if bucket := float64(l.bucket); tokens > bucket { //令牌数目不能大于桶容量
		tokens = bucket
	}

	return now, last, tokens

}

// 将token转化为所需等待时间
func (l *Limiter) durationFromTokens(tokens float64) time.Duration {
	seconds := tokens / float64(l.rate)
	return time.Nanosecond * time.Duration(1e9*seconds)
}

//计算一段时间生产的token数目
func (l *Limiter) tokensFromDuration(d time.Duration) float64 {
	sec := float64(d/time.Second) * float64(l.rate)
	nsec := float64(d%time.Second) * float64(l.rate)
	return sec + nsec/1e9
}
