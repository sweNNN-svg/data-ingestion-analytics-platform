# etl_worker/extract.py
import psycopg2
import os

def get_data():
    # Docker ortamından bağlantı bilgilerini al
    conn = psycopg2.connect(os.getenv("DATABASE_URL"))
    cur = conn.cursor()
    
    # Her şeyi getir
    cur.execute("SELECT event_type FROM raw_events")
    rows = cur.fetchall() # Örn: [('purchase',), ('click',)...]
    
    cur.close()
    conn.close()
    return rows
