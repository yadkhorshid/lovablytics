import { useEffect, useState } from 'react';

const Profile = () => {
  const [profile, setProfile] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchProfile = async () => {
      const token = localStorage.getItem('spotify_access_token');
      if (!token) {
        setError('No Spotify access token found. Please login again.');
        return;
      }

      try {
        const profileRes = await fetch('http://127.0.0.1:8080/spotify/profile', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!profileRes.ok) {
          throw new Error('Failed to fetch profile');
        }

        const profileData = await profileRes.json();
        setProfile(profileData);
      } catch (err) {
        setError(err.message);
      }
    };

    fetchProfile();
  }, []);

  if (error) return <p>Error: {error}</p>;
  if (!profile) return <p>Loading...</p>;

  return (
    <div>
      <h2>Welcome, {profile.display_name}</h2>
      <img src={profile.images?.[0]?.url} alt="Profile" width="150" />
      <p>Email: {profile.email}</p>
      <p>Followers: {profile.followers.total}</p>
      <a href={profile.external_urls.spotify} target="_blank" rel="noopener noreferrer">Open Spotify Profile</a>
    </div>
  );
};

export default Profile;
