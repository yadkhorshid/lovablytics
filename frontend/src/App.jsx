import { Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import AnalyzePage from './pages/AnalyzePage';
import SpotifyCallback from './pages/SpotifyCallback';
import Profile from './components/Profile';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/analyze" element={<AnalyzePage />} />
      <Route path="/auth/spotify/callback" element={<SpotifyCallback />} />
      <Route path="/profile" element={<Profile />} />
    </Routes>
  );
}
