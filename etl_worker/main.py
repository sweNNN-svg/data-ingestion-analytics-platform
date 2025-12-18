import time
import os
from extract import get_data
from transform import count_events
from load import save_data

def run_etl_pipeline():
    print(f"[{time.strftime('%Y-%m-%d %H:%M:%S')}] ETL Döngüsü Başlatılıyor...")
    try:
        # Ham veriyi çek (psycopg2)
        raw_rows = get_data()
        
        if not raw_rows:
            print("İşlenecek yeni veri bulunamadı. Bekleniyor...")
            return

        # Veriyi işle (Python dict aggregation)
        aggregated_data = count_events(raw_rows)
        print(f"Veri özetlendi: {aggregated_data}")

        # Analytics tablosuna yaz (Idempotent)
        save_data(aggregated_data)
        print("Analytics tablosu başarıyla güncellendi.")

    except Exception as e:
        print(f"Hata oluştu: {str(e)}")

if __name__ == "__main__":
    # Docker ayağa kalktığında DB'nin hazır olması için kısa bir bekleme
    time.sleep(5) 
    while True:
        run_etl_pipeline()
        # Her 60 saniyede bir çalış (Simulated)
        time.sleep(60)