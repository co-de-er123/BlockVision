import React, { useState } from 'react';
import axios from 'axios';

const BlockDetails = () => {
  const [blockNumber, setBlockNumber] = useState('');
  const [blockData, setBlockData] = useState(null);

  const fetchBlockDetails = async () => {
    const result = await axios.get(`http://localhost:8080/block/${blockNumber}`);
    setBlockData(result.data);
  };

  return (
    <div>
      <input 
        type="number" 
        value={blockNumber} 
        onChange={(e) => setBlockNumber(e.target.value)} 
        placeholder="Enter Block Number" 
      />
      <button onClick={fetchBlockDetails}>Get Block Details</button>
      {blockData && <pre>{JSON.stringify(blockData, null, 2)}</pre>}
    </div>
  );
};

export default BlockDetails;
