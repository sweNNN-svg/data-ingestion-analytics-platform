import psycopg2
import os

def save_data(summary):
    conn = psycopg2.connect(os.getenv("DATABASE_URL"))
    cur = conn.cursor()
    
    # ÖNCE TEMİZLE!!!! ("Basic Idempotency")
    cur.execute("DELETE FROM analytics_events")
    
    # ŞİMDİ YAZ!
    for event_name, count in summary.items():
        cur.execute(
            "INSERT INTO analytics_events (event_type, event_count, window_start) VALUES (%s, %s, NOW())",
            (event_name, count)
        )
    
    conn.commit()
    cur.close()
    conn.close()