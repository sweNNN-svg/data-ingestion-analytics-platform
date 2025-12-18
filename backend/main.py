from fastapi import FastAPI
from pydantic import BaseModel, Field
from datetime import datetime

app = FastAPI(title="Scate Ingestion API")

# Gelen paket kontrolü
class EventCreate(BaseModel):
    event_type: str = Field(..., alias="eventType")
    user_id: int = Field(..., alias="userId")
    timestamp: datetime = Field(..., alias="timestamp")
    metadata: dict = Field(default_factory=dict, alias="metadata")


# Endpoint tanımı
@app.post("/events", status_code=201)
async def ingest_event(event: EventCreate):
    # Burada event işleme yapılacak
    print(f"Received event: {event.model_dump()}")
    return event