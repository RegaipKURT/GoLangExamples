package main

import (
	"fmt"
	"math/rand"

	"../utils" //bir üst dizine çık ve utils içerisinde bul.
	"github.com/fatih/color"
	loglama "github.com/koding/logging"
	//aynı paket isimleri karışmasın diye yukarıdaki gibi isimlendirilebilir.
)

/*PAKET NEDİR?
Yukarıdaki importların hepsi birer pakettir aslında.
Github paketleri "go get github.com/isim/paket" sözdizimiyle konsoldan yüklenebilir ve bu
paketler de kullanıcılar tarafından yazılmış olan paketlerdir. Biz ise paketi indirip
import edip kendi projelerimizde kullanabiliyoruz.
Paketler bir araya toplanmış fonksiyonlar topluluğudur. Her Go programı paketlerden oluşur.
Programlar main paketinden çalışmaya başlar. Paketler projelerin düzgün tutulmasını sağlar,
İlgili fonksiyonları bir arada toplar. Derleyiciyi hızlandırır.
Örneğin: fmt'nin her kullanımda derleyici tarafından import edilmesine gerek kalmaz.
C# ve JAVA'daki namespace kavramına benzerdir.

PAKET OLUŞTURMAK İÇİN:
Bulunduğunuz dizinde libraries isimli bir klasör açıp main.go dosyanızı orada yazın.
Daha sonra libraries ile aynı dizinde utils isimli klasör açıp o klasörün altında utils.go
dosyasına yine utils paket adı vererek paket oluşturulabilir. paket "../utils" şeklinde
import edilip fonksiyonlar çağrılarak kullanılabilir. paketlerin fonksiyon isimleri
büyük harfle başlamalıdır ki başka paketlerden o fonksiyon çağrılabilsin. (private/public)
Yukarıda bahsettiğim isimler değiştirilebilir. ../paket_adı yerine gopath'e atılarak da
paket kullanmak ve import etmek mümkündür. Ayarlar yapılarak farklı şekillerde kullanılabilir.
*/
func main() {
	fmt.Println("Rastgele sayı: ", rand.Int())
	color.Red("HATA MESAJI!")
	loglama.Warning("Hata mesajı verildi.")
	loglama.Info("Kod Tamamlandı!")
	utils.Sonuc()
}
