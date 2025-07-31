import { useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';

export default function SpotifyCallback() {
  const [params] = useSearchParams();
  const navigate = useNavigate();

  useEffect(() => {
    const code = params.get('code');
    if (!code) return;

    const fetchToken = async () => {
      try {
        const res = await fetch(`http://127.0.0.1:8080/auth/spotify/callback?code=${code}`);
        const data = await res.json();

        if (data.access_token) {
          localStorage.setItem('spotify_access_token', data.access_token);
          navigate('/analyze'); 
        } else {
          console.error('Token fetch failed', data);
        }
      } catch (err) {
        console.error('Error fetching token:', err);
      }
    };

    fetchToken();
  }, [params, navigate]);

  return <p>Authenticating with Spotify...</p>;
}
