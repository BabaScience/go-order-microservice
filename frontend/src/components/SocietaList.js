// src/components/SocietaList.js

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import config from '../config';

function SocietaList() {
  const [societaList, setSocietaList] = useState([]);

  useEffect(() => {
    axios.get(`${config.API_BASE_URL}/societa`)
      .then(response => {
        const results = response.data
        setSocietaList(results);
      })
      .catch(error => console.error(error));
  }, []);

  return (
    <div>
      <h2>Societa</h2>
      {
        societaList.length > 0 ? (
          <table>
            <thead>
              <tr>
                <th>Codice Societa</th>
              </tr>
            </thead>
            <tbody>
              {societaList.map(societa => (
                <tr key={societa.id}>
                  <td>{societa.codice}</td>
                </tr>
              ))}
            </tbody>
          </table>
        ) : (
          <p>No societa found</p>
        )
      }
    </div>
  );
}

export default SocietaList;
