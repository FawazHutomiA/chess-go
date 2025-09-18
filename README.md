Siap 🚀 berikut draft **README.md** untuk project `chess-go`:

---

# 🏰 Console Chess (Go)

A simple **console-based chess game** written in **Golang**.
Supports basic chess rules, input parsing, console board display, and unit tests.

---

## 📦 Requirements

* Go **1.22+**
* Terminal that supports **ANSI colors** (for board colors)

---

## ⚙️ Setup

Clone repository lalu masuk folder project:

```bash
git clone https://github.com/FawazHutomiA/chess-go.git
cd chess-go
```

Inisialisasi modul Go:

```bash
go mod tidy
```

---

## ▶️ Run the Game

```bash
go run .
```

### Example Gameplay

```
Current player: White
    a  b  c  d  e  f  g  h
8  r  n  b  q  k  b  n  r  8
7  p  p  p  p  p  p  p  p  7
6  .  .  .  .  .  .  .  .  6
5  .  .  .  .  .  .  .  .  5
4  .  .  .  .  .  .  .  .  4
3  .  .  .  .  .  .  .  .  3
2  P  P  P  P  P  P  P  P  2
1  R  N  B  Q  K  B  N  R  1
    a  b  c  d  e  f  g  h
Enter move (e.g. b2 b3 or 1,2 2,2):
```

---

## 🎮 Move Input Format

Supports 2 formats:

1. **Algebraic** (file+rank):

   ```
   b2 b3
   e7 e5
   ```

2. **Numeric row,col** (1–8):

   ```
   1,2 2,2
   7,5 5,5
   ```

⚠️ Rows count **from top (1 = top)**, while algebraic notation `a1` is bottom-left (White’s perspective).

---

## 🧪 Run Tests

```bash
go test ./...
```

All unit tests check:

* ✅ Board initialization
* ✅ Legal & illegal moves
* ✅ Pawn behavior (1 step, 2 steps, blocked, capture)
* ✅ King capture ends the game

---

## 📌 Notes

* This is a simplified chess:

  * No **check / checkmate detection**
  * No **castling**
  * No **pawn promotion**
  * No **en passant**

Game ends when **a King is captured**.

---

## 🚀 Future Improvements

* [ ] Add check/checkmate detection
* [ ] Implement castling
* [ ] Add pawn promotion
* [ ] En passant rule
* [ ] Multiplayer over network

---

🔹 Author: Fawaz Hutomi Abdurahman
🔹 License: MIT

---