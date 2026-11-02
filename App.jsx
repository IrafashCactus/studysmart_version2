
import React, { useState, useEffect } from 'react';
import Layout from './components/Layout';
import Dashboard from './views/Dashboard';
import MindFeed from './views/MindFeed';
import AIBuddy from './views/AIBuddy';
import Quiz from './views/Quiz';
import FocusMode from './views/FocusMode';
import Leaderboard from './views/Leaderboard';
import Profile from './views/Profile';
import ParentInsights from './views/ParentInsights';
export default function App() {
  const [currentView, setCurrentView] = useState('feed');
  const [darkMode, setDarkMode] = useState(false);

  // Apply Dark Mode class to the <html> tag
  useEffect(() => {
    if (darkMode) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }, [darkMode]);

  const renderView = () => {
    switch (currentView) {
      case 'dash': return <Dashboard />;
      case 'feed': return <MindFeed />;
      case 'buddy': return <AIBuddy />;
      case 'quiz': return <Quiz />;
      case 'focus': return <FocusMode />;
      case 'leaderboard': return <Leaderboard />;
      case 'profile': return <Profile />;
      default: return <MindFeed />;
      case 'parent': return <ParentInsights />;
    }
  };

return (
  <Layout 
    currentView={currentView} 
    setView={setCurrentView} 
    darkMode={darkMode}       // <-- ARE THESE TWO 
    setDarkMode={setDarkMode} // <-- LINES HERE?
  >
    {renderView()}
  </Layout>
);
}