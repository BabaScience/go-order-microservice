// src/components/Posizione.js

import React, { useEffect, useState } from 'react';
import axios from 'axios';

import config from '../config';

function PosizioniList() {
  const [posizioni, setPosizioni] = useState([]);

  useEffect(() => {
    axios.get(`${config.API_BASE_URL}/posizione`)
      .then(response => setPosizioni(response.data))
      .catch(error => console.error(error));
  }, []);

  return (
    <div>
      <h2>Posizioni</h2>
      {
        posizioni.length > 0 ? (
          <table>
            <thead>
              <tr>
                <th>Contratto</th>
                <th>Inizio Coperatura</th>
                <th>Fine Coperatura</th>
                <th>Quantita</th>
                <th>UdM</th>
                <th>Divisione</th>
                <th>Magazzino</th>
                <th>Dilazione Pagamento</th>
                <th>Valuta</th>
                <th>Materiale</th>
                <th>Contratto Order</th>
              </tr>
            </thead>
            <tbody>
              {posizioni.map(posizione => (
                <tr key={posizione.posizione_contratto}>
                  <td>{posizione.inizio_coperatura}</td>
                  <td>{posizione.fine_coperatura}</td>
                  <td>{posizione.quantita}</td>
                  <td>{posizione.UdM}</td>
                  <td>{posizione.divisione}</td>
                  <td>{posizione.magazzino}</td>
                  <td>{posizione.dilazione_pagamento}</td>
                  <td>{posizione.valuta}</td>
                  <td>{posizione.quantita}</td>
                  <td>{posizione.materiale}</td>
                  <td>{posizione.contratto_order}</td>
                </tr>
              ))}
            </tbody>
          </table>
        ) : (
          <p>No Posizioni found</p>
        )
      }
    </div>
  );
}

export default PosizioniList;
