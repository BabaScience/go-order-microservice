import React from 'react';
import OrderList from './components/OrderList';
import SocietaList from './components/SocietaList';
import PosizioniList from './components/PosizioniList';

import './App.css';

function App() {
  return (
    <div className="App">
      <h1>Data Visualization</h1>
      <OrderList />
      <SocietaList />
      <PosizioniList />
    </div>
  );
}

export default App;
