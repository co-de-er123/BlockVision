import React from 'react';
import BlockDetails from './components/BlockDetails';
import TransactionDetails from './components/TransactionDetails';
import NetworkHealth from './components/NetworkHealth';

function App() {
  return (
    <div>
      <h1>Blockchain Explorer</h1>
      <BlockDetails />
      <TransactionDetails />
      <NetworkHealth />
    </div>
  );
}

export default App;
