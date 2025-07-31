// src/pages/LandingPage.jsx
import { useNavigate } from 'react-router-dom';
import './LandingPage.css';

export default function LandingPage() {
  const navigate = useNavigate();

  const handleLogin = () => {
    const clientId = '71cf22f55c574e398337ad23ca422f05';
    const redirectUri = 'http://127.0.0.1:5173/auth/spotify/callback';
    const scope = 'user-read-email user-top-read';

    window.location.href = `https://accounts.spotify.com/authorize?client_id=${clientId}&response_type=code&redirect_uri=${redirectUri}&scope=${scope}`;
  };

  return (
    <div className="landing-container">
      <h1>Lovablytics</h1>
      <p>Discover your musical mood and emotional energy with Spotify insights.</p>
      <button onClick={handleLogin}>Login with Spotify</button>
    </div>
  );
}
