import React from 'react';
import './TrackCard.css';

export default function TrackCard({ track }) {
  const { name, artist, genres, mood } = track;

  const moodColors = {
    happy: '#ffd700',
    upbeat: '#ff8c00',
    calm: '#87ceeb',
    energetic: '#ff4500',
    unknown: '#ccc',
  };

  const moodColor = moodColors[mood] || moodColors.unknown;

  return (
    <div className="track-card" style={{ borderColor: moodColor }}>
      <h3>{name}</h3>
      <p><em>By {artist}</em></p>
      {genres && genres.length > 0 && (
        <p className="genres">Genres: {genres.join(', ')}</p>
      )}
      <p className="mood" style={{ color: moodColor }}>
        Mood: {mood || 'Unknown'}
      </p>
    </div>
  );
}
