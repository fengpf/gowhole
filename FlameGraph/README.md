
### go pprof 工具分析 CPU Heap

  go tool pprof http://127.0.0.1:6341

  go tool pprof --pdf  http://127.0.0.1:2333/debug/pprof/profile > profile.pdf

  go tool pprof --pdf  http://127.0.0.1:2333/debug/pprof/heap > heap.pdf


### 生成火焰图
 
下载perl脚本
git clone https://github.com/brendangregg/FlameGraph

go tool pprof -seconds=10 -raw -output=a.pprof http://127.0.0.1:2333/debug/pprof/profile

./stackcollapse-go.pl a.pprof > pprof.folded  

./flamegraph.pl pprof.folded > pprof.svg