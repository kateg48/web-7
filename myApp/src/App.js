import React from 'react';
import Hello from './tasks/Hello';
import User from './tasks/Query';
import Count from './tasks/Count';
import './App.css'

const App = () => {
    return (
        <div>
            <h1>Microservices Interface</h1>
            <Hello />
            <User  />
            <Count />
        </div>
    );
};

export default App;