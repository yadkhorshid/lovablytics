import React, { useEffect, useState } from 'react';
import './AnalyzePage.css';
import TrackCard from '../components/TrackCard';

export default function AnalyzePage() {
  const [tracks, setTracks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchTopTracks = async () => {
      setLoading(true);
      setError(null);
      const token = localStorage.getItem('spotify_access_token');

      try {
        const res = await fetch('http://127.0.0.1:8080/spotify/top-tracks', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!res.ok) throw new Error('Failed to fetch top tracks');
        const data = await res.json();
        setTracks(data);
      } catch (err) {
        setError(err.message || 'Something went wrong');
      } finally {
        setLoading(false);
      }
    };

    fetchTopTracks();
  }, []);

  return (
    <div className="analyze-container">
      <h1>Your Music Mood Analysis </h1>

      {loading && <p>Loading your top tracks...</p>}
      {error && <p className="error-text">{error}</p>}

      {!loading && !error && (
        <div className="track-list">
          {tracks.map((track) => (
            <TrackCard key={track.id} track={track} />
          ))}
        </div>
      )}
    </div>
  );
}
