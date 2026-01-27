# Go Programming

คำตอบโจทย์ Go ทั้ง 5 ข้อ พร้อมคำอธิบาย

## โจทย์ที่ 1: Worker Pool (Concurrency)
Basic Concurrency: Worker Pool
โจทย์: จงเขียนฟังก์ชันชื่อ RunWorkers โดยรับ parameter 2 ตัวคือ numWorkers (จำนวนคนงาน) และ
numJobs (จำนวนงานทั้งหมด)
ให้สร้ำง Workers จำนวน numWorkers ตัวที่ทำงานพร้อมๆ กัน (Concurrent)
แต่ละ Worker จะคอยรับงานจาก Channel แล้วพิมพ์ข้อความว่า "Worker [ID] processing job [Job ID]"
ให้จำลองกำรทำงานด้วย time.Sleep เล็กน้อย(เช่น 1 วินาที)
Main function ต้องรอให้ทุกงานถูกทำ จนเสร็จสิ้นจริงๆ ก่อนถึงจะจบโปรแกรม

**โฟลเดอร์:** `1_worker_pool`

**อธิบาย:**
- ใช้ `goroutines` สร้าง worker pool ที่ทำงานพร้อมกัน
- ใช้ `channel` สำหรับส่งงานให้ workers
- ใช้ `sync.WaitGroup` เพื่อรอให้ทุกงานเสร็จสิ้น
- แต่ละ worker จะพิมพ์ข้อความและจำลองการทำงานด้วย `time.Sleep(1s)`

**วิธีรัน:**
```bash
go run 1_worker_pool/main.go
```

---

## โจทย์ที่ 2: Thread-Safe Counter
Thread-Safety: Safe Counter
โจทย์: จงสร้าง struct ชื่อ SafeCounter ที่ภายในเก็บค่า count (int) และมี Methods 2 ตัวดังน้ี:
Inc(): สำหรับเพิ่มค่า count ทีละ 1
Value(): สำหรับคืนค่า count ปัจจุบันออกมา
เงื่อนไข: โคด้น้ีต้องรองรับกรณีที่มี Goroutines จำนวนมาก(เช่น 1,000 routines) เรียกใช้Inc() พร้อมกัน
โดยที่ค่าสุดท้ายต้องถูกต้องแม่นยำ ห้ำมเกิด Race Condition

**โฟล์เดอร์:** `2_safe_counter`

**อธิบาย:**
- ใช้ `sync.Mutex` เพื่อป้องกัน race condition
- Method `Inc()` จะ lock ก่อนเพิ่มค่า แล้ว unlock
- Method `Value()` จะ lock ก่อนอ่านค่า แล้ว unlock
- ทดสอบด้วย 1,000 goroutines เพื่อยืนยันความถูกต้อง

**วิธีรัน:**
```bash
go run 2_safe_counter/main.go
```

**ทดสอบ race condition:**
```bash
go run -race 2_safe_counter/main.go
```

---

## โจทย์ที่ 3: Shape Interface
Interfaces: ระบบคำนวณพื้นที่ (Shape)
โจทย์: ให้ประกาศ(Define) Interface ชื่อ Shape ที่มี method ชื่อ Area() float64
สร้ำง Struct 2 ตัวชื่อ Rectangle (มี field Width, Height) และ Circle (มี field Radius)
Implement method Area ให้กับ Struct ทั้งสองเพื่อคำนวณพื้นที่
เขียนฟังก์ชันแยกออกมา 1 ตัว ชื่อ PrintArea(s Shape) ที่รับ Shape รูปทรงไหนก็ได้ แล้วพิมพ์ขนาดพื้นที่ออกมา

**โฟล์เดอร์:** `3_shape_interface`

**อธิบาย:**
- สร้าง `interface Shape` กับ method `Area()`
- Implement `Rectangle` และ `Circle` ที่มี method `Area()`
- ฟังก์ชัน `PrintArea()` รับ `Shape` ใดก็ได้ (polymorphism)
- Rectangle: พื้นที่ = กว้าง × สูง
- Circle: พื้นที่ = π × รัศมี²

**วิธีรัน:**
```bash
go run 3_shape_interface/main.go
```

---

## โจทย์ที่ 4: Two Sum (HashMap)
Logic & Map: หาคู่ตัวเลข(Two Sum)
โจทย์: กำหนดให้มี Slice ของตัวเลข nums := []int{2, 7, 11, 15} และค่าเป้าหมาย target := 9 
จงเขียนฟังก์ชันที่รับ nums และ target แล้วคืนค่า index ของตัวเลข 2 ตัวใน slice ที่บวกกันแล้วได้เท่ากับ
target พอดี สมมติว่ามีคำตอบที่ถูกต้องแน่นอนเพียง 1 คู่
เงื่อนไข: ห้ำมใช้ Loop ซ้อน Loop (Double for-loop) ต้องใช้วิธีที่มีประสิทธิภำพ Time Complexity ดีกว่ำ O(n^2)

**โฟล์เดอร์:** `4_two_sum`

**อธิบาย:**
- ใช้ `map` เก็บค่าที่เคยเจอและ index
- วนลูปครั้งเดียว (O(n)) แทนการวนซ้อน (O(n²))
- หา complement = target - current
- ถ้าเจอ complement ใน map แสดงว่าเจอคำตอบ

**วิธีรัน:**
```bash
go run 4_two_sum/main.go
```

---

## โจทย์ที่ 5: JSON API Server
HTTP Handler: สร้ำง JSON API ง่ายๆ
โจทย์: จงใช้ package net/http เขียน Web Server ที่รันบน port 8080 และมี endpoint ชื่อ /hello 
โดยมีเงื่อนไขดังน้ี: ต้องรับ Request Method เป็น POST เท่านั้น
รับ Body เป็น JSON: {"name": "Somchai"}
แกะค่า name ออกมาแล้วตอบกลับ (Response) เป็น JSON: {"message": "Hello Somchai"}
กรณีที่ Method ไม่ใช่ POST หรือส่ง JSON มาผิด format ให้ return HTTP Error Code ที่เหมาะสมกลับไป

**โฟล์เดอร์:** `5_json_api`

**อธิบาย:**
- สร้าง HTTP server บน port 8080
- Endpoint `/hello` รับเฉพาะ POST method
- รับ JSON `{"name": "..."}` และตอบกลับ `{"message": "Hello ..."}`
- มี error handling สำหรับ:
  - Method ไม่ใช่ POST → 405 Method Not Allowed
  - JSON ผิด format → 400 Bad Request
  - ไม่มี name field → 400 Bad Request

**วิธีรัน:**
```bash
go run 5_json_api/main.go
```

**ทดสอบด้วย curl (เปิด terminal ใหม่):**

✅ **กรณีสำเร็จ:**
```bash
curl -X POST http://localhost:8080/hello \
  -H "Content-Type: application/json" \
  -d '{"name": "Somchai"}'
```
ผลลัพธ์: `{"message":"Hello Somchai"}`

❌ **กรณี Method ผิด:**
```bash
curl -X GET http://localhost:8080/hello
```
ผลลัพธ์: `{"error":"Method not allowed. Please use POST"}`

❌ **กรณี JSON ผิด:**
```bash
curl -X POST http://localhost:8080/hello \
  -H "Content-Type: application/json" \
  -d 'invalid json'
```
ผลลัพธ์: `{"error":"Invalid JSON format"}`

---

## ความต้องการของระบบ

- Go 1.16 หรือสูงกว่า
- ไม่ต้องติดตั้ง package เพิ่มเติม (ใช้ standard library)

## หมายเหตุ

- ทุกโปรแกรมพร้อมใช้งานและมี example ในตัว
- มี error handling ครบถ้วน
- โค้ดเขียนตาม best practices ของ Go
- มี comments อธิบายส่วนสำคัญ

---

## คำสั่งรัน (เพื่อทดสอบ)

```bash
# ข้อ 1
go run 1_worker_pool/main.go

# ข้อ 2
go run 2_safe_counter/main.go

# ข้อ 3
go run 3_shape_interface/main.go

# ข้อ 4
go run 4_two_sum/main.go

# ข้อ 5 (ต้องกด Ctrl+C เพื่อหยุด server)
go run 5_json_api/main.go
```
