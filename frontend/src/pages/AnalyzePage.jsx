import React, { useState } from 'react';
import './AnalyzePage.css'; 

export default function AnalyzePage() {
  const [text, setText] = useState('');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const analyzeText = async () => {
    setLoading(true);
    setError(null);
    try {
      const res = await fetch('http://127.0.0.1:8080/analyze', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text }),
      });
      if (!res.ok) throw new Error('Server returned an error');
      const data = await res.json();
      setResult(data);
    } catch (err) {
      setError(err.message || 'Something went wrong');
      setResult(null);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app-container">
      <h1>Lovablytics</h1>

      <textarea
        rows="5"
        placeholder="Enter some text to analyze..."
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button onClick={analyzeText} disabled={loading || !text.trim()}>
        {loading ? 'Analyzing...' : 'Analyze'}
      </button>

      {error && <p className="error-text">{error}</p>}
      {result && (
        <div className="results-box">
          <h3>Results:</h3>
          <pre>{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}

      <div style={{ marginTop: '2rem' }}>
        <a
          href={`https://accounts.spotify.com/authorize?client_id=71cf22f55c574e398337ad23ca422f05&response_type=code&redirect_uri=http://127.0.0.1:5173/auth/spotify/callback&scope=user-read-email`}
        >
          <button>Login with Spotify</button>
        </a>
      </div>
    </div>
  );
}
