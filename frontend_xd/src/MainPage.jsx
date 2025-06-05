import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './MainPage.css';

function MainPage() {
  const [activities, setActivities] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get('http://localhost:8000/actividades')
      .then((response) => {
        setActivities(response.data);
        setError(null);
      })
      .catch(() => {
        setError('No se pudieron cargar las actividades');
      });
  }, []);

  return (
    <div className="main-container">
      <h1>Actividades</h1>
      {error && <p className="error">{error}</p>}
      <div className="cards">
        {activities.map((act) => (
          <div key={act.id} className="card">
            <img src={act.imagen} alt={act.titulo} />
            <h2>{act.titulo}</h2>
            <p>{act.descripcion}</p>
            <p>
              <strong>Dia:</strong> {act.dia} - <strong>Horario:</strong> {act.horario}
            </p>
            <p>
              <strong>Cupo:</strong> {act.cupo}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default MainPage;
