# Event Tracking & Analytics Platform

## 1. Proje Özeti (Project Overview)
Bu proje, yüksek hacimli event verilerinin (ingestion) karşılanması, işlenmesi (ETL) ve son kullanıcıya sunulması (Dashboard) süreçlerini kapsayan uçtan uca (End-to-End) bir veri mühendisliği simülasyonudur. Amaç, sadece veri akışını sağlamak değil, aynı zamanda ölçeklenebilir, hataya dayanıklı (fault-tolerant) ve modern bir mikroservis mimarisi inşa etmektir.

**Sistem Akışı:**
1. **Ingestion:** Veri, `FastAPI` tarafından yüksek performansla karşılanır ve `PostgreSQL` veritabanındaki "Raw" katmanına yazılır.
2. **Processing:** Bağımsız çalışan bir Python `ETL Worker`, ham veriyi belirli zaman pencerelerinde işleyerek analitik tablolara dönüştürür.
3. **Visualization:** Son olarak, `Next.js` tabanlı dashboard, bu işlenmiş veriyi gerçek zamanlıya yakın bir hızda görselleştirir.

---
---

## 2. Mimari Kararlar ve Ödünler (Key Engineering Decisions & Trade-offs)
Bir sistem tasarlarken yapılan tercihler, kodun kendisi kadar önemlidir. Bu projede alınan kritik kararlar şunlardır:

### 2.1. Servis Ayrıştırması (Decoupling)
ETL işlemlerini (Transform/Load) doğrudan API'nin içine gömmek yerine, bağımsız bir **Worker** servisine ayırdım.

* **Neden?** API'nin (Ingestion Layer) tek sorumluluğu, gelen trafiği en düşük gecikmeyle (latency) karşılamaktır. Eğer ağır analitik sorguları API üzerinde çalıştırsaydık, yüksek trafik anında sistem darboğaza (bottleneck) girerdi. Worker servisinin ayrılması, "Ingestion" ve "Processing" katmanlarının bağımsız olarak ölçeklenebilmesine (**Horizontal Scaling**) olanak tanır.

### 2.2. Veri Bütünlüğü ve Idempotency
ETL sürecinde **"DELETE -> INSERT"** mantığını benimsedim.

* **Neden?** Dağıtık sistemlerde ağ hataları veya servis kesintileri nedeniyle bir işlem yarım kalabilir veya tekrar tetiklenebilir. Eğer sadece INSERT yapsaydık, aynı zaman aralığı için veriyi iki kez işleyip (duplicate data) raporları bozabilirdik. Tasarladığım yapı, ETL scripti aynı saat aralığı için 100 kez de çalışsa sonucun her zaman aynı ve doğru olmasını (**Idempotency**) garanti eder.

### 2.3. Konteynerizasyon (Containerization)
Proje; Backend, Database, ETL Worker ve Frontend olmak üzere 4 farklı Docker servisi olarak kurgulandı.

* **Neden?** "Benim makinemde çalışıyordu" sorununu ortadan kaldırmak ve ortam bağımsız (**portable**) bir yapı kurmak önceliğimdi. Her servis kendi bağımlılıklarıyla izole edilmiştir, bu da CI/CD süreçlerine tam uyumluluk sağlar.

---

## 3. Karşılaşılan Zorluklar ve Çözümler (Challenges & Solutions)
Geliştirme sürecinde karşılaşılan problemler ve uyguladığım çözüm stratejileri:

### 3.1. Race Conditions (Servis Başlatma Sırası)
* **Sorun:** `docker-compose up` komutu verildiğinde tüm servisler aynı anda ayağa kalkmaya çalışıyor. Ancak ETL servisi, veritabanı (PostgreSQL) henüz bağlantıları kabul etmeye hazır olmadan bağlanmaya çalıştığında çöküyordu.
* **Çözüm:** Docker Compose dosyasında `depends_on` ve healthcheck mekanizmalarını kullandım. Ayrıca, servislerin içinde veritabanı bağlantısı kurulana kadar bekleyen (**retry logic**) dirençli bir yapı kurguladım. Bu, sistemin kendi kendini iyileştirmesini (**resilience**) sağladı.

### 3.2. CORS ve Network Routing
* **Sorun:** Backend ve Frontend farklı konteynerlerde çalışırken, tarayıcı (Browser) üzerinden yapılan isteklerde CORS (Cross-Origin Resource Sharing) hataları ve ağ erişim sorunları yaşandı.
* **Çözüm:**
    * Backend tarafında `CORSMiddleware` kullanarak, Frontend'den gelen isteklere güvenli bir şekilde izin verdim.
    * Frontend'in "Server-Side" ve "Client-Side" isteklerini ayrıştırdım. Statik IP ve port yönlendirmeleriyle, tarayıcının konteyner içindeki API'ye sorunsuz erişmesini sağladım.

---

## 4. Teknoloji Yığını (Tech Stack)
Proje, modern ve endüstri standardı teknolojiler seçilerek geliştirilmiştir:

* **Backend (Ingestion):** Python, FastAPI (High concurrency için Async/Await), Pydantic.
* **Database (Storage):** PostgreSQL (Raw logs ve Analytics summary tabloları).
* **ETL (Processing):** Python, Pandas & Raw SQL (Batch processing).
* **Frontend (Visualization):** Next.js.
* **Infrastructure:** Docker & Docker Compose.

---

## 5. Kurulum (How to Run)
Tüm sistemi tek bir komutla ayağa kaldırabilirsiniz. Docker'ın sisteminizde kurulu olması yeterlidir.

1. Projeyi klonlayın ve ana dizine gidin.
2. Aşağıdaki komutu çalıştırın:

```bash
docker-compose up --build
