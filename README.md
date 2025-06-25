# Worker-Pool-URL-Pinger

This project demonstrates a **concurrent worker pool** in Go to efficiently **ping multiple URLs** using HTTP HEAD requests.

Instead of launching hundreds of goroutines, it uses a **fixed number of workers** to process jobs from a shared queue (channel), making the program more efficient and scalable.

---

##  Features

-  Worker Pool pattern using goroutines & channels  
-  HTTP HEAD request to check URL availability  
-  Graceful handling of valid and invalid URLs  
-  WaitGroup synchronization  
-  Buffered channels for job distribution and result collection  
-  Fully concurrent and efficient

---

##  Concepts Practiced

- Goroutines  
- Channels (job queue & result stream)  
- WaitGroups  
- Concurrency control (worker pool)  
- Error handling with timeouts

---

##  How It Works

1. **Jobs** (URLs) are sent into a `jobs` channel.  
2. **Workers** (fixed number of goroutines) pull from this channel.  
3. Each worker pings the URL and sends the result into a `results` channel.  
4. The main thread collects and prints the results after all work is done.

---

##  Run the Project

```bash
go run main.go
