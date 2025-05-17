# 📦 ServeAPI — Mock JSON API สำหรับ Frontend

ServeAPI คือเว็บแอปและ REST API ที่ออกแบบมาเพื่อให้ **Frontend สามารถทดสอบกับ API ปลอม (Mock API)** ได้ในขณะรอการพัฒนา Backend จริง

ช่วยให้คุณสามารถ:
- 📤 สร้างข้อมูล JSON ปลอม (POST)
- 📋 ดึงข้อมูลทั้งหมดกลับมา (GET)
- ❌ ลบข้อมูล JSON บางชุด (DELETE)
- 💻 ใช้หน้าเว็บ GUI ในการส่งหรือคัดลอก API link ได้ง่าย

---

## 🧩 ใช้ทำอะไร?

- ให้ทีม Frontend ใช้ API ปลอมระหว่างรอ Backend
- ทดสอบ UI ที่ต้องใช้ข้อมูล JSON จริง
- แชร์ลิงก์ API ให้เพื่อนร่วมทีมทดสอบ UI
- ใช้งานแบบ local หรือในเครือข่ายทีม

---

## 🚀 วิธีใช้งาน

### POST JSON
ส่ง JSON เข้าระบบเพื่อสร้าง API ปลอม

```http
POST /postapi
## ▶️ วิธีรันระบบ

1. ติดตั้ง Go modules และ dependencies:

```bash
go mod tidy
go run .

```

## IMAGE
![image](https://github.com/user-attachments/assets/3e377ac3-f34d-4047-8245-22320cb05a8b)
![image](https://github.com/user-attachments/assets/6861a533-fbba-4bf4-87b7-d51da4bfe660)
