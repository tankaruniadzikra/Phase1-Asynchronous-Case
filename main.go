// sebagai developer sering mengalami warning
// buat program sederhana dengan implementasi goroutine yang mana mengirim pesan log secara langsung dengan asyincronus
// output:
// application continues after logging...
// [INFO] - User 123 logged in
// [WARN] - Memory user is high
// [ERROR] - Failed to fetch data from API
// Main application finished

package main

import (
	"fmt"
	"time"
)

// 1. Define Struct
type LogMessage struct {
	Level   string
	Message string
}

// 2. Fungsi Asinkron dengan dua parameter, yaitu level dan message.
func levelAsync(level string, message string) {
	time.Sleep(2 * time.Second) // Untuk menunda eksekusi selama 2 detik
	fmt.Printf("[%s] - %s\n", level, message)
}

func main() {
	// 3. Define Dataset
	LogMessages := []LogMessage{ // Deklarasi dan inisialisasi variabel LogMessages. Variabel ini adalah sebuah slice yang berisi elemen-elemen dari tipe data LogMessage.
		{Level: "INFO", Message: "User 123 logged in"},
		{Level: "WARN", Message: "Memory user is high"},
		{Level: "ERROR", Message: "Failed to fetch data from API"},
	}

	// 4. Eksekusi Asinkron
	for _, LogMessage := range LogMessages {
		// Penggunaan go di depan, fungsi ini akan dilakukan secara asinkron dalam goroutine terpisah.
		// Hal ini memungkinkan program untuk melanjutkan eksekusi tanpa harus menunggu setiap pemanggilan levelAsync selesai. Dengan kata lain, pesan log akan dikirim secara asinkron dan program utama dapat melanjutkan eksekusinya sendiri.
		go levelAsync(LogMessage.Level, LogMessage.Message)
	}

	// 5. Informasi Status
	fmt.Println("Application continues after logging...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main application finished")

}
