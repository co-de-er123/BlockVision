import React, { useEffect, useState } from 'react';
import axios from 'axios';

const NetworkHealth = () => {
  const [networkHealth, setNetworkHealth] = useState(null);

  useEffect(() => {
    const fetchNetworkHealth = async () => {
      const result = await axios.get('http://localhost:8080/network-health');
      setNetworkHealth(result.data);
    };
    fetchNetworkHealth();
  }, []);

  return (
    <div>
      <h3>Network Health</h3>
      {networkHealth && <pre>{JSON.stringify(networkHealth, null, 2)}</pre>}
    </div>
  );
};

export default NetworkHealth;
