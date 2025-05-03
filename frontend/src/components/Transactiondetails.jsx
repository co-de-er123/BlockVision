import React, { useState } from 'react';
import axios from 'axios';

const TransactionDetails = () => {
  const [txHash, setTxHash] = useState('');
  const [txData, setTxData] = useState(null);

  const fetchTransactionDetails = async () => {
    const result = await axios.get(`http://localhost:8080/transaction/${txHash}`);
    setTxData(result.data);
  };

  return (
    <div>
      <input 
        type="text" 
        value={txHash} 
        onChange={(e) => setTxHash(e.target.value)} 
        placeholder="Enter Transaction Hash" 
      />
      <button onClick={fetchTransactionDetails}>Get Transaction Details</button>
      {txData && <pre>{JSON.stringify(txData, null, 2)}</pre>}
    </div>
  );
};

export default TransactionDetails;
