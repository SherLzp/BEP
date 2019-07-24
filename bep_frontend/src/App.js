import React from 'react';
import './App.css';
import { HashRouter, Router, Link } from 'react-router-dom';
import 'semantic-ui-css/semantic.min.css'
import HomepageLayout from './components/HomePageLayout';

function App() {
    return (
        <HashRouter>
            <HomepageLayout />
        </HashRouter>
    );
}

export default App;
