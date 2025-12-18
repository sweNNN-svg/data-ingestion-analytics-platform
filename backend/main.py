from fastapi import FastAPI
from pydantic import BaseModel, Field
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlalchemy.orm import sessionmaker
from models import Base, RawEvent
from fastapi import Depends, HTTPException
from fastapi.middleware.cors import CORSMiddleware

# veritabanı bağlantısı
DATABASE_URL = "postgresql+asyncpg://user:password@db/events_db"

# Asenkron veritabanı motoru ol
engine = create_async_engine(DATABASE_URL, echo=True)
AsyncSessionLocal = sessionmaker(
    engine,
    class_=AsyncSession,
    expire_on_commit=False
)

# FastAPI uygulaması oluşturma
app = FastAPI(title="Scate Ingestion API")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Her yerden gelen isteğe izin ver
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Gelen paket kontrolü
class EventCreate(BaseModel):
    event_type: str = Field(..., alias="eventType")
    user_id: int = Field(..., alias="userId")
    timestamp: datetime = Field(..., alias="timeStamp")
    metadata: dict = Field(default_factory=dict, alias="metaData")

# Veritabanı oturumu
async def get_db():
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)

    db = AsyncSessionLocal()
    try:
        yield db
    finally:
        await db.close()


# Endpoint tanımı
@app.post("/events", status_code=201)
async def ingest_event(event: EventCreate, db: AsyncSession = Depends(get_db)):
    """
    Bu fonksiyon, gelen olay verilerini alır ve veritabanına kaydeder.
    async def olduğu için veritabanı işlemleri asenkron olarak gerçekleştirilir.
    201 kodu ile başarılı bir şekilde oluşturulduğunu belirtir.
    """
    try:
        new_event = RawEvent(
            event_type=event.event_type,
            user_id=event.user_id,
            timestamp=event.timestamp,
            metadata_field=event.metadata
        )
        # Veritabanına ekle ve işlemi onayla
        db.add(new_event)
        await db.commit()
        return {"message": "Event ingested successfully"}
    
    except Exception as e:
        # Hata durumunda işlemi geri al
        await db.rollback()
        raise HTTPException(status_code=500, detail=str(e))
    


@app.get("/analytics/summary")
async def get_analytics_summary(db=Depends(get_db)):
    from sqlalchemy import text
    try:
        # SQL ile direkt analytics tablosundan çekiyoruz
        query = text("SELECT event_type, event_count, window_start FROM analytics_events")
        result = await db.execute(query)
        rows = result.fetchall()
        
        # Veriyi liste formatına çevirip dönüyoruz
        return [
            {
                "event_type": r[0], 
                "event_count": r[1], 
                "window_start": r[2].isoformat() if r[2] else None
            } 
            for r in rows
        ]
    except Exception as e:
        print(f"Veri çekme hatası: {e}")
        return []