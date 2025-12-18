from sqlalchemy import Column, Integer, String, DateTime, func
from sqlalchemy.dialects.postgresql import JSONB
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class RawEvent(Base):
    __tablename__ = 'raw_events'

    id = Column(Integer, primary_key=True, index=True)
    event_type = Column(String, index=True)
    user_id = Column(Integer, index=True)
    timestamp = Column(DateTime(timezone=True)) # Zaman damgası için timezone ekledim
    metadata_field = Column(JSONB)
    created_at = Column(DateTime, default=func.now())


class AnalyticsEvent(Base):
    __tablename__ = 'analytics_events'

    id = Column(Integer, primary_key=True, index=True)
    event_type = Column(String)
    event_count = Column(Integer)
    window_start = Column(DateTime(timezone=True))

    