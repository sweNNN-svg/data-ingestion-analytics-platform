"use client";
import { useEffect, useState } from "react";

export default function AdminDashboard() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  // Backend'den veriyi çeken fonksiyon
  const fetchAnalytics = async () => {
    try {
      const response = await fetch("http://192.168.100.2:8000/analytics/summary");
      const result = await response.json();
      setData(result);
    } catch (error) {
      console.error("Veri çekilemedi:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAnalytics();
    const interval = setInterval(fetchAnalytics, 30000); // 30 saniyede bir güncelle
    return () => clearInterval(interval);
  }, []);

  return (
    <main style={{ padding: '2rem', backgroundColor: '#f8fafc', minHeight: '100vh', fontFamily: 'sans-serif' }}>
      <div style={{ maxWidth: '800px', margin: '0 auto', backgroundColor: 'white', padding: '2rem', borderRadius: '12px', boxShadow: '0 4px 6px -1px rgb(0 0 0 / 0.1)' }}>
        
        <h1 style={{ color: '#1e293b', marginBottom: '1.5rem', borderBottom: '2px solid #e2e8f0', paddingBottom: '0.5rem' }}>
          📊 Scate Analytics Dashboard
        </h1>

        {loading ? (
          <p>Yükleniyor...</p>
        ) : (
          <table style={{ width: '100%', borderCollapse: 'collapse', marginTop: '1rem' }}>
            <thead>
              <tr style={{ backgroundColor: '#f1f5f9', textAlign: 'left' }}>
                <th style={{ padding: '12px', color: '#475569' }}>Event Type</th>
                <th style={{ padding: '12px', color: '#475569' }}>Total Count</th>
                <th style={{ padding: '12px', color: '#475569' }}>Last Window</th>
              </tr>
            </thead>
            <tbody>
              {data.map((item, idx) => (
                <tr key={idx} style={{ borderBottom: '1px solid #f1f5f9' }}>
                  <td style={{ padding: '12px', fontWeight: 'bold', color: '#334155' }}>{item.event_type.toUpperCase()}</td>
                  <td style={{ padding: '12px' }}>
                    <span style={{ backgroundColor: '#dcfce7', color: '#166534', padding: '4px 12px', borderRadius: '20px', fontSize: '0.9rem' }}>
                      {item.event_count} events
                    </span>
                  </td>
                  <td style={{ padding: '12px', color: '#64748b', fontSize: '0.85rem' }}>
                    {new Date(item.window_start).toLocaleString()}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
        
        {data.length === 0 && !loading && (
          <p style={{ textAlign: 'center', color: '#94a3b8', marginTop: '2rem' }}>Henüz işlenmiş veri bulunamadı.</p>
        )}
      </div>
    </main>
  );
}