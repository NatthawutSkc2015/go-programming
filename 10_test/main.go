package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ════════════════════════════════════════════
// 💀 BUG #1: Goroutine Leak — ผีที่ไม่ยอมตาย
// ════════════════════════════════════════════
// goroutine ถูกสร้างแต่ไม่มีทางออก
// memory พุ่งขึ้นเรื่อย ๆ จนระบบหมดแรง
func goroutineLeak() {
	for i := 0; ; i++ {
		go func(id int) {
			ch := make(chan struct{}) // channel ที่ไม่มีใคร send ให้เลย
			<-ch                      // block ตลอดกาล → goroutine ค้างในอากาศ
		}(i)
		time.Sleep(time.Millisecond)
	}
}

// ════════════════════════════════════════════
// 💀 BUG #2: Deadlock — งูกินหาง
// ════════════════════════════════════════════
// Goroutine A รอ Goroutine B
// Goroutine B รอ Goroutine A
// → ระบบแช่แข็งสมบูรณ์ Go runtime ตรวจเจอแล้ว panic: all goroutines are asleep
func deadlock() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		time.Sleep(10 * time.Millisecond) // เปิดโอกาสให้ goroutine อื่น lock mu2 ก่อน
		mu2.Lock()                        // รอ mu2 ← แต่ goroutine ข้างล่างถือ mu2 แล้ว
		defer mu1.Unlock()
		defer mu2.Unlock()
	}()

	go func() {
		mu2.Lock()
		time.Sleep(10 * time.Millisecond)
		mu1.Lock() // รอ mu1 ← แต่ goroutine ข้างบนถือ mu1 แล้ว
		defer mu2.Unlock()
		defer mu1.Unlock()
	}()
}

// ════════════════════════════════════════════
// 💀 BUG #3: Memory Leak — น้ำรั่วในถัง
// ════════════════════════════════════════════
// cache ที่ไม่มีวันล้าง → RAM พุ่งขึ้นไม่หยุด
// บั๊กนี้อันตรายเพราะ "ไม่พังทันที" แต่ค่อย ๆ กินจนตาย
var cache = make(map[string][]byte)
var cacheMu sync.Mutex

func memoryLeak() {
	for i := 0; ; i++ {
		key := fmt.Sprintf("key-%d", i)
		data := make([]byte, 1024*1024) // 1 MB ต่อ entry
		cacheMu.Lock()
		cache[key] = data // เพิ่มทุกครั้ง ไม่เคย delete
		cacheMu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

// ════════════════════════════════════════════
// 💀 BUG #4: Stack Overflow — ก้นบึ้งที่ไม่มี
// ════════════════════════════════════════════
// Infinite recursion → stack โตจนระเบิด
// Go จะ panic: stack overflow
func infiniteRecursion(n int) int {
	return infiniteRecursion(n + 1) // ไม่มี base case → เรียกตัวเองไปเรื่อย ๆ
}

// ════════════════════════════════════════════
// 💀 BUG #5: Data Race — สองคนแก้ไฟล์เดียวกัน
// ════════════════════════════════════════════
// goroutine หลายตัว อ่าน/เขียน variable เดียวกันพร้อมกัน
// ผล: ข้อมูลเสีย, crash แบบสุ่ม, บั๊กที่ reproduce ไม่ได้
// ตรวจได้ด้วย: go run -race main.go
var counter int // ← ไม่มี mutex ป้องกัน!

func dataRace() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // READ + WRITE พร้อมกัน = undefined behavior
		}()
	}
	wg.Wait()
	fmt.Println("counter (corrupted):", counter) // ไม่เท่า 1000 แน่นอน
}

// ════════════════════════════════════════════
// 💀 BUG #6: Nil Pointer Dereference — ผีร้ายแห่ง Runtime
// ════════════════════════════════════════════
type Config struct {
	DB *Database
}

type Database struct {
	Host string
}

func nilPointerDereference() {
	var cfg *Config          // nil pointer!
	fmt.Println(cfg.DB.Host) // panic: nil pointer dereference
	// เกิดบ่อยมากเวลา error handling ไม่ครบ
}

// ════════════════════════════════════════════
// 💀 BUG #7: Channel Deadlock — ส่งจดหมายไม่มีคนรับ
// ════════════════════════════════════════════
func unbufferedChannelBlock() {
	ch := make(chan int) // unbuffered channel
	ch <- 42             // บล็อกตรงนี้ทันที เพราะไม่มีใครรอรับ
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

// ════════════════════════════════════════════
// 💀 BUG #8: Goroutine + Closure Trap — ค่าที่หายไป
// ════════════════════════════════════════════
// Classic Go gotcha — goroutine capture variable by reference
func closureTrap() {
	results := make([]func(), 5)
	for i := 0; i < 5; i++ {
		results[i] = func() {
			fmt.Println(i) // ทุก goroutine print ค่า i สุดท้าย (= 5) หมด!
		}
	}
	for _, f := range results {
		f()
	}
	// Output: 5 5 5 5 5  ← ไม่ใช่ 0 1 2 3 4
}

// ════════════════════════════════════════════
// แสดงสถานะ Goroutine ที่รั่ว
// ════════════════════════════════════════════
func printStats() {
	for {
		fmt.Printf("🔴 Goroutines: %d | Memory: %.1f MB\n",
			runtime.NumGoroutine(),
			float64(memStats())/1024/1024,
		)
		time.Sleep(time.Second)
	}
}

func memStats() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	fmt.Println("=== Bug Zoo ===")
	fmt.Println("เลือก bug ที่ต้องการจำลอง:")
	fmt.Println("1) Goroutine Leak  2) Deadlock  3) Memory Leak")
	fmt.Println("4) Stack Overflow  5) Data Race 6) Nil Pointer")
	fmt.Println()

	// --- เปิดตัวเดียวตามต้องการ ---

	// go goroutineLeak()
	// go printStats()
	// time.Sleep(time.Minute)

	// deadlock()
	// time.Sleep(time.Second)

	// go memoryLeak()
	// go printStats()
	// time.Sleep(time.Minute)

	// infiniteRecursion(0)

	dataRace() // safe demo — ดู race: go run -race main.go

	// nilPointerDereference()

	// unbufferedChannelBlock()

	// closureTrap()
}
