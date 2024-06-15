import './App.css';
import React from 'react';
import { Link, Route, Routes, BrowserRouter as Router} from "react-router-dom";

import ScrapPage from './pages/scrapPage';

function App() {
  return (
    <div>
      <Router>
        <Routes>
          <Route path='/' element={<ScrapPage />}/>
        </Routes>
      </Router>
    </div>
  );
}

export default App;