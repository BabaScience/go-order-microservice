// src/components/OrderList.js

import React, { useEffect, useState } from 'react';
import axios from 'axios';

import config from '../config';

function OrderList() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    axios.get(`${config.API_BASE_URL}/orders`)
      .then(response => {
        const results = response.data
        setOrders(results);
      })
      .catch(error => console.error(error));
  }, []);

  return (
    <div>
      <h2>Orders</h2>
      {
        orders.length > 0 ? (
          <table>
            <thead>
              <tr>
                <th>Contratto</th>
                <th>Committente</th>
                <th>Data Creazione</th>
                <th>Data Inizio Validita</th>
                <th>Data Fine Validita</th>
                <th>Codici Chiave</th>
                <th>Matricola Creazione</th>
                <th>Codice Societa</th>
              </tr>
            </thead>
            <tbody>
              {orders.map(order => (
                <tr key={order.id}>
                  <td>{order.contratto}</td>
                  <td>{order.committente}</td>
                  <td>{order.data_creazione}</td>
                  <td>{new Date(order.data_inizio_validita).toDateString() || order.data_inizio_validita}</td>
                  <td>{new Date(order.data_fine_validita).toDateString() || order.data_fine_validita }</td>
                  <td>{order.codice_chiave}</td>
                  <td>{order.matricola_creazione}</td>
                  <td>{order.codice_societa}</td>
                </tr>
              ))}
            </tbody>
          </table>
        ) : (
          <p>No orders found</p>
        )
      }
    </div>
  );
}

export default OrderList;
